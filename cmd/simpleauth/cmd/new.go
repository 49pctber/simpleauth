package cmd

import (
	"fmt"
	"log"

	"github.com/49pctber/simpleauth"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new user",
	Long:  `Create a new user`,
	Run: func(cmd *cobra.Command, args []string) {

		var username, password, password2 string
		var err error

		fname, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error parsing file flag: %v\n", err)
		}
		simpleauth.Configure(fname)

		admin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			log.Fatalf("error parsing admin flag: %v\n", err)
		}

		if admin {
			fmt.Println("** CREATING ADMINISTRATIVE USER **")
		}

		fmt.Printf("Username: ")
		fmt.Scanln(&username)
		if !simpleauth.ValidateUsername(username) {
			log.Fatal("invalid username")
		}

		u := simpleauth.FindUser(username)
		if u != nil {
			log.Fatal("user already exists")
		}

		fmt.Printf("Password: ")
		fmt.Scanln(&password)
		if len(password) < 8 {
			log.Fatal("password too short")
		}

		fmt.Printf("Verify Password: ")
		fmt.Scanln(&password2)

		if password != password2 {
			log.Fatal("passwords don't match")
		}

		err = simpleauth.AddUser(username, password, admin)
		if err != nil {
			log.Fatalf("error creating user: %v\n", err)
		}
	},
}

func init() {
	userCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("admin", "a", false, "indicates whether a user should be given the admin permission")
}
