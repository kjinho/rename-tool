/*
Package cmd for extension

Copyright © 2021 Jin-Ho King <j@kingesq.us>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// extensionCmd represents the extension command
var extensionCmd = &cobra.Command{
	Use:   "extension [new extension] [files ...]",
	Short: "Replaces the file extensions",
	Long: `
Replaces the three letter file extensions with a new
three letter extension, or appends the new three letter
extension when the original does not have a three letter
extension.
`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("extension called")
		for _, file := range args[1:] {
			newFilename := genNewFilename(file, args[0])
			err := os.Rename(file, newFilename)
			if err != nil {
				log.Fatalf("`extension` error renaming file `%s` to `%s`\nError: %s", file, newFilename, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(extensionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extensionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extensionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
