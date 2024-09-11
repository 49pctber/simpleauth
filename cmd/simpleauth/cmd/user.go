package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "list all users",
	Long:  `list all users`,
	// Run: func(cmd *cobra.Command, args []string) {	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
