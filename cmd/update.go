package cmd

import (
	"fmt"

	"github.com/nednella/bootstrap.sh/internal/ui"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update binary and pull latest content",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("update")
		fmt.Println("update: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
