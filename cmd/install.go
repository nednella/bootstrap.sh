package cmd

import (
	"fmt"

	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages from Brewfile",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("install")
		fmt.Println("install: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
