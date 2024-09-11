package cmd

import (
	"log"
	"os"

	"github.com/49pctber/simpleauth"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a simpleauth configuration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fname, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error parsing file flag: %v\n", err)
		}

		_, err = os.Stat(fname)
		if err == nil {
			log.Fatalf("%v already exists", fname)
		}

		err = simpleauth.NewAuthConfig(fname)
		if err != nil {
			log.Fatalf("error creating new configuration file: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
