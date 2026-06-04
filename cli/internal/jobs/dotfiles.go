package jobs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

const backupTimestampFormat = "20060102-150405"

func Dotfiles() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	dotfilesDir := filepath.Join(cfg.InstallPath, "dotfiles")
	if !utils.Exists(dotfilesDir) {
		return fmt.Errorf("dotfiles directory not found: %s", dotfilesDir)
	}

	backupDir := filepath.Join(cfg.BackupPath, time.Now().Format(backupTimestampFormat))

	programs, err := os.ReadDir(dotfilesDir)
	if err != nil {
		return err
	}

	for _, program := range programs {
		if !program.IsDir() {
			continue
		}
		err := linkProgram(filepath.Join(dotfilesDir, program.Name()), program.Name(), backupDir)
		if err != nil {
			return err
		}
	}

	ui.Success("Dotfiles linked")
	return nil
}

func linkProgram(programDir, programName, backupDir string) error {
	entries, err := os.ReadDir(programDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		src := filepath.Join(programDir, entry.Name())
		target := resolveTarget(entry.Name(), programName)
		err := link(src, target, backupDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// resolveTarget applies the convention: dot-prefix → $HOME, else → $XDG_CONFIG_HOME/<program>/
func resolveTarget(filename, programName string) string {
	if strings.HasPrefix(filename, ".") {
		return filepath.Join(utils.Home, filename)
	}
	return filepath.Join(utils.ConfigDir, programName, filename)
}

func link(src, target, backupDir string) error {
	name := utils.DisplayName(target)

	switch {
	case !utils.Exists(target):
		return createLink(src, target, name)
	case utils.IsSymlinkedTo(target, src):
		ui.Info(name + " — already linked")
		return nil
	case utils.IsSymlinked(target):
		return replaceLink(src, target, name)
	default:
		return backupAndLink(src, target, backupDir, name)
	}
}

func createLink(src, target, name string) error {
	err := utils.MkdirAll(filepath.Dir(target), 0755)
	if err != nil {
		return err
	}
	err = utils.Symlink(src, target)
	if err != nil {
		return err
	}
	ui.Info(name + " — linked")
	return nil
}

func replaceLink(src, target, name string) error {
	current, err := os.Readlink(target)
	if err != nil {
		return err
	}
	err = utils.Remove(target)
	if err != nil {
		return err
	}
	err = utils.Symlink(src, target)
	if err != nil {
		return err
	}
	ui.Warn(name + " — re-linked (was → " + current + ")")
	return nil
}

func backupAndLink(src, target, backupDir, name string) error {
	dest := filepath.Join(backupDir, name)
	err := utils.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return err
	}
	err = utils.Rename(target, dest)
	if err != nil {
		return err
	}
	ui.Warn(name + " — backed up to " + backupDir)
	err = utils.Symlink(src, target)
	if err != nil {
		return err
	}
	ui.Info(name + " — linked")
	return nil
}
