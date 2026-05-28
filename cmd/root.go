package cmd

import (
	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a fresh macOS install",
	Long: `Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

Installs Homebrew + Brewfile, symlinks dotfiles into $HOME, and applies macOS preferences.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.Banner()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
