package config

import (
	_ "embed"

	"github.com/nednella/bootstrap.sh/internal/utils"
	"gopkg.in/yaml.v3"
)

// Preflight needs this config file to know where to clone in the first place,
// so it must be available before any clone exists, hence we use embedded.

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

	cfg.InstallPath = utils.ExpandHome(cfg.InstallPath)
	cfg.BackupPath = utils.ExpandHome(cfg.BackupPath)
	return cfg, nil
}
