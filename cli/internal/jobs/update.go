package jobs

import (
	"encoding/json"
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

func Update() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	latest, err := latestVersion(cfg.RepoURL)
	if err != nil {
		return err
	}

	if semver.Compare(internal.Version, latest) >= 0 {
		ui.Info("Already up to date (" + internal.Version + ")")
		return nil
	}

	ui.Info("Updating binary " + internal.Version + " → " + latest + " (requires sudo) ...")
	url := cfg.RepoURL + "/releases/latest/download/" + binaryAsset
	staged := binaryDest + ".new"

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
	err = utils.Command("sudo", "mv", staged, binaryDest)
	if err != nil {
		return err
	}

	ui.Success("Updated → " + latest)
	return nil
}

func latestVersion(repoURL string) (string, error) {
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
