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
	if path == "~" {
		return Home
	}
	if !strings.HasPrefix(path, "~/") {
		return path
	}
	return filepath.Join(Home, path[2:])
}

// ExpandPath resolves a manifest destination: the `~` prefix plus the
// `$HOME` / `$XDG_CONFIG_HOME` tokens. $XDG_CONFIG_HOME routes through
// ConfigDir, so it honours the env var and falls back to ~/.config — never empty.
func ExpandPath(path string) string {
	expanded := os.Expand(path, func(key string) string {
		switch key {
		case "HOME":
			return Home
		case "XDG_CONFIG_HOME":
			return ConfigDir
		default:
			return os.Getenv(key)
		}
	})
	return ExpandHome(expanded)
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
