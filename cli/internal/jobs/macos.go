package jobs

import (
	"slices"

	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

type macosDefault struct {
	currentHost bool // some settings must target the current user instead of global
	domain      string
	key         string
	valueType   string
	value       string
}

var dockSettings = []macosDefault{}

var finderSettings = []macosDefault{}

var menuBarSettings = []macosDefault{
	{true, "com.apple.controlcenter", "BatteryShowPercentage", "-bool", "true"},
}

var services = []string{"ControlCenter", "Dock", "Finder", "SystemUIServer"}

func MacOS() error {
	closeSystemSettings()

	settings := slices.Concat(dockSettings, finderSettings, menuBarSettings)
	for _, s := range settings {
		err := applyDefault(s)
		if err != nil {
			return err
		}
	}

	restartServices()
	ui.Success("macOS defaults applied")
	return nil
}

// System Settings can overwrite our writes when it closes, so quit it first.
func closeSystemSettings() {
	ui.Info("Closing System Settings ...")
	_ = utils.Command("pkill", "-x", "System Settings")
}

func applyDefault(s macosDefault) error {
	ui.Info(s.key + " → " + s.value)
	args := []string{"write", s.domain, s.key, s.valueType, s.value}
	if s.currentHost {
		args = append([]string{"-currentHost"}, args...)
	}
	return utils.Command("defaults", args...)
}

func restartServices() {
	ui.Info("Restarting services ...")
	for _, app := range services {
		_ = utils.Command("killall", app)
	}
}
