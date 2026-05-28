package utils

import (
	"os"
	"os/exec"
	"strings"

	"github.com/nednella/bootstrap.sh/internal/ui"
)

var DryRun bool

func Command(name string, args ...string) error {
	if DryRun {
		ui.Dry(name + " " + strings.Join(args, " "))
		return nil
	}
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
