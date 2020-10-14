package cmd

import (
	"os"
    "io/ioutil"
	"github.com/spf13/cobra"
	"fmt"
)

var valid_languages    = map[string]bool   { "python": true, }
var valid_license      = map[string]bool   { "apache": true, "gnugpl3": true, "mit": true, }
var license_conversion = map[string]string { "apache": Apache_str, "gnugpl3": GNUGPL3_str, "mit": MIT_str, }

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
		if name != "" { create_project(name) } else { panic("Invalid name") }
		if language != "" {
			if valid_languages[language] { create_language_env(language) } else { fmt.Println("Language unrecognised, use 'orwell patchfix -l' to remedy") }
		}
		if license != "" {
			if valid_license[license] { create_license_file(license) } else { fmt.Println("License unrecognised, use 'orwell patchfix -x' to remedy") }
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

func create_project(name string) {
	err := os.Mkdir(name, 0755); check(err)
	os.Chdir(name)
	wd, _ := os.Getwd()
    create_readme(wd, name)
}

func create_readme(directory string, name string) {
   new_file, err := os.Create("README.md"); check(err)
   file_path := directory + "/README.md"
   readme_content := ("# " + name + readme)
   file_contents := []byte(readme_content)
   err = ioutil.WriteFile(file_path, file_contents, 0644); check(err)
   new_file.Close()
   fmt.Printf("README.md Generated\n")
}

func create_language_env(language string) {

}

func create_license_file(license_type string) {
   new_file, err := os.Create("LICENSE.md"); check(err)
   wd, _ := os.Getwd()
   file_path := wd + "/LICENSE.md"
   err = ioutil.WriteFile(file_path, []byte(license_conversion[license_type]), 0644); check(err)
   new_file.Close()
   fmt.Printf("LICENSE.md Generated\n")
}

func connect_repository(repository string) {

}

func check(e error) { if e != nil { panic(e) } }
