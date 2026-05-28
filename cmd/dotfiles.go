package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Symlink dotfiles into $HOME / XDG",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dotfiles: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(dotfilesCmd)
}
