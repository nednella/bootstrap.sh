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

	ui.Info("Pulling latest changes ...")
	err = utils.Command("git", "-C", cfg.InstallPath, "pull", "--rebase", "--autostash")
	if err != nil {
		return err
	}

	ui.Success("Synced")
	return nil
}
