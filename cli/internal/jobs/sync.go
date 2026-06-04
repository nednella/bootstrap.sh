package jobs

import (
	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

func Sync() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	clone := cfg.InstallPath
	dirty, err := isDirty(clone)
	if err != nil {
		return err
	}

	if dirty {
		ui.Info("Stashing local changes ...")
		err = utils.Command("git", "-C", clone, "stash", "push", "-u")
		if err != nil {
			return err
		}
	}

	ui.Info("Pulling latest changes ...")
	err = utils.Command("git", "-C", clone, "pull")
	if err != nil {
		return err
	}

	if dirty {
		ui.Info("Restoring local changes ...")
		err = utils.Command("git", "-C", clone, "stash", "pop")
		if err != nil {
			return err
		}
	}

	ui.Success("Synced")
	return nil
}

func isDirty(clone string) (bool, error) {
	if utils.DryRun {
		return false, nil
	}
	out, err := utils.Output("git", "-C", clone, "status", "--porcelain")
	return out != "", err
}
