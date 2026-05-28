package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update binary and pull latest content",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
