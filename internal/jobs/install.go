package jobs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

const (
	brewBin       = "/opt/homebrew/bin/brew"
	brewBinDir    = "/opt/homebrew/bin"
	brewInstaller = `NONINTERACTIVE=1 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
)

func Install() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	err = ensureHomebrew()
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

func ensureHomebrew() error {
	if utils.Exists(brewBin) {
		ui.Info("Homebrew — already installed")
	} else {
		ui.Info("Homebrew — installing (non-interactive)")
		err := utils.Command("/bin/bash", "-c", brewInstaller)
		if err != nil {
			return err
		}
	}
	return addBrewToPath()
}

func addBrewToPath() error {
	path := os.Getenv("PATH")
	if strings.Contains(path, brewBinDir) {
		return nil
	}
	return os.Setenv("PATH", brewBinDir+":"+path)
}
