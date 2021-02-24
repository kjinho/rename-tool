/*
Package cmd contains what's needed to implement rename-tool

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
	"strconv"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move [format string] [destination directory] [start number] [stop number]",
	Short: "Batch move files numbered files into a directory",
	Long: `
The move command batch moves numbered files in the current 
working directory to a given subdirectory. The format string 
provides the naming convention for the files to be moved. 
Sprintf is used to generate the wildcard match for the format
string using numbers between the start and stop numbers 
provided.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("move called with following arguments")
		for i, item := range args {
			fmt.Printf("%3d | %s\n", i, item)
		}
		startno, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalf("Value for `startno` is not an integer: `%s`\nError: %s", args[2], err)
		}
		stopno, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			log.Fatalf("Value for `stopno` is not an integer: `%s`\nError: %s", args[3], err)
		}
		batchMove(args[0], args[1], startno, stopno)
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
