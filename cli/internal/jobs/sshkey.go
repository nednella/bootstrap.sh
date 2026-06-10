package jobs

import (
	"fmt"
	"path/filepath"

	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/nednella/bootstrap.sh/internal/utils"
)

func SSHKey() error {
	sshDir := filepath.Join(utils.Home, ".ssh")
	keyPath := filepath.Join(sshDir, "id_ed25519")
	pubPath := keyPath + ".pub"

	if utils.Exists(keyPath) {
		ui.Info("Key already exists — " + utils.DisplayName(keyPath))
	} else {
		err := generateKey(sshDir, keyPath)
		if err != nil {
			return err
		}
		ui.Success("Key generated")
	}

	err := utils.Command("bash", "-c", "pbcopy < '"+pubPath+"'")
	if err != nil {
		return err
	}

	ui.Success("Public key copied to clipboard")
	return nil
}

func generateKey(sshDir, keyPath string) error {
	email, err := gitEmail()
	if err != nil {
		return err
	}

	err = utils.MkdirAll(sshDir, 0700)
	if err != nil {
		return err
	}

	// ssh-keygen prompts for the passphrase interactively; an empty one leaves
	// the private key unencrypted on disk.
	return utils.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", keyPath)
}

func gitEmail() (string, error) {
	email, err := utils.Output("git", "config", "user.email")
	if err != nil || email == "" {
		return "", fmt.Errorf("git user.email is not set — configure it before generating an SSH key")
	}
	return email, nil
}
