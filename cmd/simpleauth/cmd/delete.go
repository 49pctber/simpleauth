package cmd

import (
	"log"

	"github.com/49pctber/simpleauth"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [user]",
	Short: "Delete a user",
	Long:  `Delete a user`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fname, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error parsing file flag: %v\n", err)
		}
		simpleauth.Configure(fname)

		username := args[0]

		err = simpleauth.DeleteUser(username)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	userCmd.AddCommand(deleteCmd)
}
