package jobs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

const (
	brewBin       = "/opt/homebrew/bin/brew"
	brewBinDir    = "/opt/homebrew/bin"
	brewInstaller = `NONINTERACTIVE=1 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
)

func Preflight() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	brewMissing := !utils.Exists(brewBin)
	gitMissing := !utils.Lookup("git")
	repoMissing := !utils.Exists(filepath.Join(cfg.InstallPath, ".git"))

	if brewMissing || gitMissing || repoMissing {
		ui.Header("preflight")

		err = ensureHomebrew(brewMissing)
		if err != nil {
			return err
		}
		err = ensureGit(gitMissing)
		if err != nil {
			return err
		}
		err = ensureRepo(cfg, repoMissing)
		if err != nil {
			return err
		}
	}

	return addBrewToPath()
}

func ensureHomebrew(missing bool) error {
	if !missing {
		return nil
	}
	ui.Info("Homebrew — installing (non-interactive) ...")
	return utils.Command("/bin/bash", "-c", brewInstaller)
}

func ensureGit(missing bool) error {
	if !missing {
		return nil
	}
	return fmt.Errorf("git not found — install the Xcode Command Line Tools (xcode-select --install)")
}

func ensureRepo(cfg *config.Config, missing bool) error {
	if !missing {
		return nil
	}
	ui.Info("Repository — cloning to " + cfg.InstallPath + " ...")
	return utils.Command("git", "clone", cfg.RepoURL, cfg.InstallPath)
}

func addBrewToPath() error {
	if utils.DryRun {
		return nil
	}
	path := os.Getenv("PATH")
	if pathHasDir(path, brewBinDir) {
		return nil
	}
	return os.Setenv("PATH", brewBinDir+":"+path)
}

func pathHasDir(path, dir string) bool {
	for _, p := range filepath.SplitList(path) {
		if p == dir {
			return true
		}
	}
	return false
}
