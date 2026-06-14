package jobs

import (
	"fmt"
	"os"
	"path/filepath"
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

	destinations, err := loadManifest(dotfilesDir)
	if err != nil {
		return err
	}

	programs, err := os.ReadDir(dotfilesDir)
	if err != nil {
		return err
	}

	backupDir := filepath.Join(cfg.BackupPath, time.Now().Format(backupTimestampFormat))

	for _, program := range programs {
		if !program.IsDir() {
			continue
		}
		dest, err := destination(destinations, program.Name())
		if err != nil {
			return err
		}
		err = linkProgram(filepath.Join(dotfilesDir, program.Name()), dest, backupDir)
		if err != nil {
			return err
		}
	}

	ui.Success("Dotfiles linked")
	return nil
}

func DotfilesUndo() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	dotfilesDir := filepath.Join(cfg.InstallPath, "dotfiles")
	if !utils.Exists(dotfilesDir) {
		return fmt.Errorf("dotfiles directory not found: %s", dotfilesDir)
	}

	destinations, err := loadManifest(dotfilesDir)
	if err != nil {
		return err
	}

	programs, err := os.ReadDir(dotfilesDir)
	if err != nil {
		return err
	}

	backupDir := latestBackup(cfg.BackupPath)

	for _, program := range programs {
		if !program.IsDir() {
			continue
		}
		dest, err := destination(destinations, program.Name())
		if err != nil {
			return err
		}
		err = unlinkProgram(filepath.Join(dotfilesDir, program.Name()), dest, backupDir)
		if err != nil {
			return err
		}
	}

	ui.Success("Dotfiles unlinked")
	return nil
}

type manifest struct {
	Destinations map[string]string `yaml:"destinations"`
}

// read dotfiles.yaml to determine where to link each program
func loadManifest(dotfilesDir string) (map[string]string, error) {
	manifest := manifest{}
	err := utils.LoadYAML(filepath.Join(dotfilesDir, "dotfiles.yaml"), &manifest)
	if err != nil {
		return nil, err
	}
	return manifest.Destinations, nil
}

func destination(destinations map[string]string, programName string) (string, error) {
	dest, ok := destinations[programName]
	if !ok {
		return "", fmt.Errorf("no destination declared for %q in dotfiles.yaml", programName)
	}
	return utils.ExpandPath(dest), nil
}

func linkProgram(programDir, dest, backupDir string) error {
	entries, err := os.ReadDir(programDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		src := filepath.Join(programDir, entry.Name())
		target := filepath.Join(dest, entry.Name())
		err := link(src, target, backupDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func unlinkProgram(programDir, dest, backupDir string) error {
	entries, err := os.ReadDir(programDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		src := filepath.Join(programDir, entry.Name())
		target := filepath.Join(dest, entry.Name())
		err := unlink(src, target, backupDir)
		if err != nil {
			return err
		}
	}
	return nil
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

// return the latest backup directory, or "" if none
func latestBackup(backupPath string) string {
	entries, err := os.ReadDir(backupPath)
	if err != nil {
		return ""
	}

	latest := ""
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() > latest {
			latest = entry.Name()
		}
	}
	if latest == "" {
		return ""
	}
	return filepath.Join(backupPath, latest)
}

func unlink(src, target, backupDir string) error {
	name := utils.DisplayName(target)

	if !utils.IsSymlinkedTo(target, src) {
		ui.Info(name + " — not linked")
		return nil
	}

	err := utils.Remove(target)
	if err != nil {
		return err
	}
	ui.Info(name + " — unlinked")

	backup := filepath.Join(backupDir, name)
	if backupDir == "" || !utils.Exists(backup) {
		return nil
	}

	err = utils.Rename(backup, target)
	if err != nil {
		return err
	}
	ui.Info(name + " — restored from " + backupDir)
	return nil
}
