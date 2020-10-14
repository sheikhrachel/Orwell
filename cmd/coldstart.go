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
	"os"
    "io/ioutil"
	"github.com/spf13/cobra"
	"fmt"
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
		//fmt.Println(Apache_str)
		if name != "" {
			create_project(name)
		} else { panic("Invalid name") }
		if language != "" {
			
		}
		if license != "" {
			valid_license := map[string]bool {
			    "apache": true,
			    "gnugpl3": true,
			    "mit": true,
			}
			if valid_license[license] {
				create_license_file(license)
			} else { fmt.Println("License unrecognised, use 'orwell patchfix -x' to remedy")}
		}
		if repository != "" {
			
		}
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

func create_project(name string) {
	err := os.Mkdir(name, 0755); check(err)
	os.Chdir(name)
	wd, _ := os.Getwd()
    create_readme(wd, name)

    // s, _ := os.Getwd()
    // fmt.Println(s)
}

func create_readme(directory string, name string) {
   new_file, err := os.Create("README.md"); check(err)
   file_path := directory + "/README.md"
   readme_content := ("# " + name + readme)
   file_contents := []byte(readme_content)
   err = ioutil.WriteFile(file_path, file_contents, 0644); check(err)
   new_file.Close()
}

func create_license_file(license_type string) {

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
