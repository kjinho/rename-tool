/*
Package cmd zip
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

	"github.com/spf13/cobra"
)

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip [folders ...]",
	Short: "Creates zip archives for provided folders",
	Long: `
Creates zip archives for the provided folders, using the
folder names with a ".zip" extension for the zip file names.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zip called")
		for _, folder := range args {
			zipfoldername := genZipFilename(folder)
			err := zipFolder(folder, zipfoldername)
			if err != nil {
				log.Fatalf("`zip` error in compressing `%s` to `%s`\nError: %s", folder, zipfoldername, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
