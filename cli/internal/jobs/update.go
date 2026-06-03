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

	err = updateBinary(cfg.RepoURL)
	if err != nil {
		return err
	}
	err = updateContent(cfg.InstallPath)
	if err != nil {
		return err
	}

	ui.Success("Updated")
	return nil
}

func updateBinary(repoURL string) error {
	latest, err := latestVersion(repoURL)
	if err != nil {
		return err
	}

	if semver.Compare(internal.Version, latest) >= 0 {
		ui.Info("Binary — up to date (" + internal.Version + ")")
		return nil
	}

	ui.Info("Binary — updating " + internal.Version + " → " + latest + " (sudo) ...")
	url := repoURL + "/releases/latest/download/" + binaryAsset
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
	return utils.Command("sudo", "mv", staged, binaryDest)
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

func updateContent(clone string) error {
	dirty, err := isDirty(clone)
	if err != nil {
		return err
	}

	if dirty {
		ui.Info("Content — stashing local changes ...")
		err = utils.Command("git", "-C", clone, "stash", "push", "-u")
		if err != nil {
			return err
		}
	}

	ui.Info("Content — pulling latest ...")
	err = utils.Command("git", "-C", clone, "pull")
	if err != nil {
		return err
	}

	if dirty {
		ui.Info("Content — restoring local changes ...")
		return utils.Command("git", "-C", clone, "stash", "pop")
	}
	return nil
}

func isDirty(clone string) (bool, error) {
	if utils.DryRun {
		return false, nil
	}
	out, err := utils.Output("git", "-C", clone, "status", "--porcelain")
	return out != "", err
}
