package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nednella/bootstrap.sh/internal/ui"
)

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

func Lookup(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func MkdirAll(path string, perm os.FileMode) error {
	if DryRun {
		ui.Dry("mkdir -p " + path)
		return nil
	}
	return os.MkdirAll(path, perm)
}

func Output(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	return strings.TrimSpace(string(out)), err
}

func PromptSudo() error {
	if DryRun {
		ui.Dry("sudo -v")
		return nil
	}
	return Command("sudo", "-v", "-p", ui.SudoPrompt())
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
