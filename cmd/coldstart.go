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

// coldstartCmd represents the coldstart command
var coldstartCmd = &cobra.Command{
	Use:   "coldstart",
	Short: "Project Generation",
	Long: `Coldstart takes in a number of arguments to rapidly spin up a new codebase and environment to work in.
Reference the flags below for the necessary options

Supported languages: python
Supported licenses: apache, gnugpl3, mit
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		license, _ := cmd.Flags().GetString("license")
		repository, _ := cmd.Flags().GetString("repository")
		fmt.Println(name)
		fmt.Println(language)
		fmt.Println(license)
		fmt.Println(repository)
		fmt.Println("coldstart ended")
	},
}

func init() {
	rootCmd.AddCommand(coldstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//coldstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	coldstartCmd.Flags().StringP("name", "n", "", "Set the project name")
	coldstartCmd.Flags().StringP("language", "l", "", "Set the codebase language")
	coldstartCmd.Flags().StringP("license", "x", "", "Set the license")
	coldstartCmd.Flags().StringP("repository", "g", "", "Set the repository link")
}
