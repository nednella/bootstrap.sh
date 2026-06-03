package jobs

import (
	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
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
	url := repoURL + "/releases/latest/download/" + binaryAsset
	staged := binaryDest + ".new"

	ui.Info("Binary — updating " + binaryDest + " (sudo) ...")
	err := utils.PromptSudo()
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
