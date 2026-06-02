package utils

import (
	"os"
	"path/filepath"
)

var (
	Home      = mustHome()
	ConfigDir = resolveConfigDir()
)

func DisplayName(target string) string {
	rel, err := filepath.Rel(Home, target)
	if err != nil {
		return target
	}
	return rel
}

func Exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

func IsSymlinked(path string) bool {
	info, err := os.Lstat(path)
	return err == nil && info.Mode()&os.ModeSymlink != 0
}

func IsSymlinkedTo(target, src string) bool {
	current, err := os.Readlink(target)
	return err == nil && current == src
}

func mustHome() string {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return h
}

func resolveConfigDir() string {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		return xdg
	}
	return filepath.Join(Home, ".config")
}
