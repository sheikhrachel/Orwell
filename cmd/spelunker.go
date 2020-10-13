/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// spelunkerCmd represents the spelunker command
var spelunkerCmd = &cobra.Command{
	Use:   "spelunker",
	Short: "Test generation",
	Long: `Spelunker rapidly parses through an existing codebase and generates a test directory (if missing) and necessary test files and modules
Reference the flags below for the necessary options

Supported languages: python
Supported libraries: requests`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spelunker called")
	},
}

func init() {
	rootCmd.AddCommand(spelunkerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spelunkerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spelunkerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
