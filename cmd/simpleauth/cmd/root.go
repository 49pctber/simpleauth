/*
Copyright © 2024 49pctber

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/49pctber/simpleauth"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simpleauth",
	Short: "",
	Long:  ``,
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", simpleauth.DefaultConfigFilename, "path to configuration file")
}
