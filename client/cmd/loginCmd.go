package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "command to login to an existant account",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
