package utils

import (
	"os"
	"path/filepath"
)

var (
	Home      = mustHome()
	ConfigDir = resolveConfigDir()
)

func mustHome() string {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return h
}

// resolveConfigDir returns $XDG_CONFIG_HOME if set, else $HOME/.config.
func resolveConfigDir() string {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		return xdg
	}
	return filepath.Join(Home, ".config")
}

// Exists reports whether anything exists at path (symlink, file, or dir).
func Exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

// DisplayName returns target as a $HOME-relative path for log output.
func DisplayName(target string) string {
	rel, err := filepath.Rel(Home, target)
	if err != nil {
		return target
	}
	return rel
}

// IsSymlinked reports whether path is a symlink (without following it).
func IsSymlinked(path string) bool {
	info, err := os.Lstat(path)
	return err == nil && info.Mode()&os.ModeSymlink != 0
}

// IsSymlinkedTo reports whether target is a symlink pointing at src.
func IsSymlinkedTo(target, src string) bool {
	current, err := os.Readlink(target)
	return err == nil && current == src
}
