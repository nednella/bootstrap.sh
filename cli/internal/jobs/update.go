package jobs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nednella/bootstrap.sh/internal"
	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
	"golang.org/x/mod/semver"
)

const (
	binaryDest  = "/usr/local/bin/bootstrap"
	binaryAsset = "bootstrap-darwin-arm64"
)

func Update() (err error) {
	if internal.Version == "dev" {
		ui.Info("Skipping self-update (development build)")
		return nil
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	latest, err := getLatestRelease(cfg.RepoURL)
	if err != nil {
		return err
	}

	if !semver.IsValid(latest) {
		return fmt.Errorf("unexpected release tag: %s", latest)
	}

	if semver.Compare(internal.Version, latest) >= 0 {
		ui.Info("Already up to date (" + internal.Version + ")")
		return nil
	}

	ui.Info("Updating binary " + internal.Version + " → " + latest + " (requires sudo) ...")
	url := cfg.RepoURL + "/releases/latest/download/" + binaryAsset
	err = replaceBinary(url, binaryDest)
	if err != nil {
		return err
	}

	ui.Success("Updated → " + latest)
	return nil
}

func UpdateTag(tag string) error {
	if !semver.IsValid(tag) {
		return fmt.Errorf("invalid release tag: %s", tag)
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	downgrade := semver.IsValid(internal.Version) && semver.Compare(tag, internal.Version) < 0
	if downgrade && !utils.DryRun {
		ok := utils.Confirm("Downgrade " + internal.Version + " → " + tag + "?")
		if !ok {
			ui.Info("Cancelled")
			return nil
		}
	}

	ui.Info("Installing binary " + tag + " (requires sudo) ...")
	url := cfg.RepoURL + "/releases/download/" + tag + "/" + binaryAsset
	err = replaceBinary(url, binaryDest)
	if err != nil {
		return err
	}

	ui.Success("Installed → " + tag)
	return nil
}

func UpdateList() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	releases, err := getAvailableReleases(cfg.RepoURL)
	if err != nil {
		return err
	}
	if len(releases) == 0 {
		ui.Info("No releases found")
		return nil
	}

	for _, release := range releases {
		if release == internal.Version {
			ui.Info(release + " (current)")
			continue
		}
		ui.Info(release)
	}
	return nil
}

func getAvailableReleases(repoURL string) ([]string, error) {
	api := strings.Replace(repoURL, "https://github.com/", "https://api.github.com/repos/", 1) + "/releases"
	out, err := utils.Output("curl", "-fsSL", api)
	if err != nil {
		return nil, err
	}

	var releases []struct {
		TagName string `json:"tag_name"`
	}
	err = json.Unmarshal([]byte(out), &releases)
	if err != nil {
		return nil, err
	}

	tags := make([]string, len(releases))
	for i, release := range releases {
		tags[i] = release.TagName
	}
	return tags, nil
}

func getLatestRelease(repoURL string) (string, error) {
	api := strings.Replace(repoURL, "https://github.com/", "https://api.github.com/repos/", 1) + "/releases/latest"
	out, err := utils.Output("curl", "-fsSL", api)
	if err != nil {
		return "", err
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	err = json.Unmarshal([]byte(out), &release)
	return release.TagName, err
}

// replaceBinary downloads url to a staged path and atomically moves it over dest,
// removing the staged file if any step fails. Every step needs sudo because dest
// lives in a system directory.
func replaceBinary(url, dest string) (err error) {
	staged := dest + ".new"
	defer func() {
		if err != nil {
			_ = utils.Command("sudo", "rm", "-f", staged)
		}
	}()

	err = utils.PromptSudo()
	if err != nil {
		return err
	}
	err = utils.Command("sudo", "curl", "-fL", url, "-o", staged)
	if err != nil {
		return err
	}
	err = utils.Command("sudo", "chmod", "+x", staged)
	if err != nil {
		return err
	}
	_ = utils.Command("sudo", "xattr", "-d", "com.apple.quarantine", staged)
	return utils.Command("sudo", "mv", staged, dest)
}
