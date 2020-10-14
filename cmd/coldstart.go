package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var coldstartCmd = &cobra.Command{
	Use:   "coldstart",
	Short: "Project Generation",
	Long:  `Coldstart takes in a number of arguments to rapidly spin up a new codebase and environment to work in.
Reference the flags below for the necessary options

Supported languages: python
Supported licenses: apache, gnugpl3, mit
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _       := cmd.Flags().GetString("name")
		language, _   := cmd.Flags().GetString("language")
		license, _    := cmd.Flags().GetString("license")
		repository, _ := cmd.Flags().GetString("repository")
		directory, _  := os.Getwd()
		if name != "" { create_project(directory, name) } else { panic("Invalid name") }
		if language != "" {
			if valid_languages[language] { create_language_env(directory, language) } else { fmt.Println("Language unrecognised, use 'orwell patchfix -l' to remedy") }
		}
		if license != "" {
			if valid_license[license] { create_license_file(directory, license) } else { fmt.Println("License unrecognised, use 'orwell patchfix -x' to remedy") }
		}
		if repository != "" { connect_repository(repository) }
		fmt.Printf("Project [%s] Created!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(coldstartCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//coldstartCmd.PersistentFlags().String("foo", "", "A help for foo")
	coldstartCmd.Flags().StringP("name", "n", "", "Set the project name")
	coldstartCmd.Flags().StringP("language", "l", "", "Set the codebase language")
	coldstartCmd.Flags().StringP("license", "x", "", "Set the license")
	coldstartCmd.Flags().StringP("repository", "g", "", "Set the repository link")
}