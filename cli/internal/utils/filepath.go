package utils

import (
	"os"
	"path/filepath"
	"strings"
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

func ExpandHome(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}
	return filepath.Join(Home, path[2:])
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
