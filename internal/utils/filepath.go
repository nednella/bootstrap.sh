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
