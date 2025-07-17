package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "command to create an account",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
