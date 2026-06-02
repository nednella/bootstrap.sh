package jobs

import (
	"fmt"
	"path/filepath"

	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

func Install() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	brewfile := filepath.Join(cfg.InstallPath, "Brewfile")
	if !utils.Exists(brewfile) {
		return fmt.Errorf("Brewfile not found: %s", brewfile)
	}

	ui.Info("Installing packages from " + brewfile)
	err = utils.Command("brew", "bundle", "--file="+brewfile)
	if err != nil {
		return err
	}

	ui.Success("Packages installed")
	return nil
}
