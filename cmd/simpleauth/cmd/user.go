package cmd

import (
	"fmt"
	"log"

	"github.com/49pctber/simpleauth"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Manage users`,
	Run: func(cmd *cobra.Command, args []string) {
		fname, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error parsing file flag: %v\n", err)
		}
		simpleauth.Configure(fname)

		usernames := simpleauth.GetUsernames()

		fmt.Printf("Information for %s:\n", fname)
		fmt.Printf("  Number of Existing Users: %v\n", len(usernames))
		fmt.Printf("  Usernames: %v\n", usernames)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
