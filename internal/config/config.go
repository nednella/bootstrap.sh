package config

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed default_config.yaml
var defaultConfigYAML []byte

type Config struct {
	InstallPath string `yaml:"install_path"`
	BackupPath  string `yaml:"backup_path"`
	RepoURL     string `yaml:"repo_url"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := yaml.Unmarshal(defaultConfigYAML, cfg)
	if err != nil {
		return nil, err
	}
	
	cfg.InstallPath = expandHome(cfg.InstallPath)
	cfg.BackupPath = expandHome(cfg.BackupPath)
	return cfg, nil
}

func expandHome(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	return filepath.Join(home, path[2:])
}
