package jobs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nednella/bootstrap.sh/internal/config"
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

type macosDefault struct {
	Domain      string `yaml:"domain"`
	Key         string `yaml:"key"`
	ValueType   string `yaml:"value_type"`
	Value       string `yaml:"value"`
	CurrentHost bool   `yaml:"current_host"`
}

type macosSettings struct {
	Defaults []macosDefault `yaml:"defaults"`
	Services []string       `yaml:"services"`
}

func MacOS() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	path := filepath.Join(cfg.InstallPath, "macos", "settings.yaml")
	if !utils.Exists(path) {
		return fmt.Errorf("macOS settings not found: %s", path)
	}

	settings := &macosSettings{}
	err = utils.LoadYAML(path, settings)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", path, err)
	}

	closeSystemSettings()

	for _, d := range settings.Defaults {
		err := applyDefault(d)
		if err != nil {
			return err
		}
	}

	restartServices(settings.Services)
	ui.Success("macOS settings applied")
	return nil
}

// System Settings can overwrite our writes when it closes, so quit it first.
func closeSystemSettings() {
	ui.Info("Closing System Settings ...")
	_ = utils.Command("pkill", "-x", "System Settings")
}

func applyDefault(d macosDefault) error {
	if d.Domain == "" || d.Key == "" || d.ValueType == "" {
		return fmt.Errorf("incomplete macOS setting, domain/key/value_type all required: %+v", d)
	}
	ui.Info(d.Domain + " " + d.Key + " → " + d.Value)
	args := []string{"write", d.Domain, d.Key, "-" + d.ValueType, d.Value}
	if d.CurrentHost {
		args = append([]string{"-currentHost"}, args...)
	}
	return utils.Command("defaults", args...)
}

func restartServices(services []string) {
	ui.Info("Restarting services (" + strings.Join(services, ", ") + ") ...")
	for _, app := range services {
		_ = utils.Command("killall", app)
	}
}
