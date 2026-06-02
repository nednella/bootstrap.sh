package utils

import (
	"fmt"
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

func MkdirAll(path string, perm os.FileMode) error {
	if DryRun {
		ui.Dry("mkdir -p " + path)
		return nil
	}
	return os.MkdirAll(path, perm)
}

func Remove(path string) error {
	if DryRun {
		ui.Dry("rm " + path)
		return nil
	}
	return os.Remove(path)
}

func Rename(oldpath, newpath string) error {
	if DryRun {
		ui.Dry(fmt.Sprintf("mv %s %s", oldpath, newpath))
		return nil
	}
	return os.Rename(oldpath, newpath)
}

func Symlink(src, dst string) error {
	if DryRun {
		ui.Dry(fmt.Sprintf("ln -s %s %s", src, dst))
		return nil
	}
	return os.Symlink(src, dst)
}
