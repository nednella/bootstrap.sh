package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a fresh macOS install",
	Long: `Bootstrap a fresh Mac quicker than the time it takes to make a cuppa.

Installs Homebrew + Brewfile, symlinks dotfiles into $HOME, and applies macOS preferences.`,
}

func Execute() error {
	return rootCmd.Execute()
}
