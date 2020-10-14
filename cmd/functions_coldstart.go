package cmd

import (
	"fmt"
	"os"
	"os/exec"
    "io/ioutil"
	"net/url"
)

var valid_languages    = map[string]bool   { "python": true, }
var valid_license      = map[string]bool   { "apache": true, "gnugpl3": true, "mit": true, }
var license_conversion = map[string]string { "apache": Apache_str, "gnugpl3": GNUGPL3_str, "mit": MIT_str, }

func create_project(directory string, name string) {
	err   := os.Mkdir(name, 0755); check(err); os.Chdir(name)
	directory, _ = os.Getwd()
    create_readme(directory, name)
}

func create_readme(directory string, name string) {
   new_file, err  := os.Create("README.md"); check(err)
   file_path      := directory + "/README.md"
   readme_content := ("# " + name + readme)
   file_contents  := []byte(readme_content)
   err             = ioutil.WriteFile(file_path, file_contents, 0644); check(err)
   new_file.Close()
   fmt.Println("README.md Generated")
}

func create_license_file(directory string, license_type string) {
   new_file, err := os.Create("LICENSE.md"); check(err)
   err            = ioutil.WriteFile(directory + "/LICENSE.md", []byte(license_conversion[license_type]), 0644); check(err)
   new_file.Close()
   fmt.Println("LICENSE.md Generated")
}

func create_language_env(directory string, language string) {
	if language == "python" { python_setup(directory) }
}

func python_setup(directory string) {
	python_env_files(directory)
	python_src_files()
	fmt.Println("Python Environment Generated")
}

func python_env_files(directory string) {
   new_file, err := os.Create(".gitignore"); check(err)
   err            = ioutil.WriteFile(directory + "/.gitignore", []byte(gitignore_str), 0644); check(err); new_file.Close()
   new_file, err  = os.Create("Pipfile"); check(err)
   err            = ioutil.WriteFile(directory + "/Pipfile", []byte(pipfile_str), 0644); check(err); new_file.Close()
   new_file, err  = os.Create(".pre-commit-config.yaml"); check(err)
   err            = ioutil.WriteFile(directory + "/.pre-commit-config.yaml", []byte(precommit_str), 0644); check(err); new_file.Close()
   new_file, err  = os.Create("tox.ini"); check(err)
   err            = ioutil.WriteFile(directory + "/tox.ini", []byte(tox_str), 0644); check(err); new_file.Close()
   new_file, err  = os.Create(".travis.yml"); check(err)
   err            = ioutil.WriteFile(directory + "/.travis.yml", []byte(travis_str), 0644); check(err); new_file.Close()
}

func python_src_files() {
	err           := os.Mkdir("src", 0755); check(err); os.Chdir("src")
	directory, _  := os.Getwd()
	new_file, err := os.Create(".gitignore"); check(err)
    err            = ioutil.WriteFile(directory + "/.gitignore", []byte(gitignoresrc_str), 0644); check(err); new_file.Close()
    new_file, err  = os.Create("api_key.py"); check(err); new_file.Close()
    new_file, err  = os.Create("main.py"); check(err); new_file.Close(); os.Chdir("..")
    
	err            = os.Mkdir("tests", 0755); check(err); os.Chdir("tests")
	directory, _   = os.Getwd()
	new_file, err  = os.Create("test_main.py"); check(err)
    err            = ioutil.WriteFile(directory + "/test_main.py", []byte("import pytest"), 0644); check(err); new_file.Close(); os.Chdir("..")
}

func connect_repository(repository string) {
	if isValidUrl(repository){
		cmd := exec.Command("git", "init"); err := cmd.Run(); check(err); fmt.Println("initialised git")
		cmd  = exec.Command("git", "add", "."); err = cmd.Run(); check(err); fmt.Println("repository staged")
		cmd  = exec.Command("git", "commit", "-m", "'init"); err = cmd.Run(); check(err); fmt.Println("repository committed")
		cmd  = exec.Command("git", "remote", "add", "origin", repository); err = cmd.Run(); check(err); fmt.Println("git remote configured")
		cmd  = exec.Command("git", "push", "-u", "origin", "master"); err = cmd.Run(); check(err); fmt.Println("git repo pushed")
	} else { fmt.Println("Repository URL unrecognised, use 'orwell patchfix -g' to remedy")}
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil { return false }

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" { return false }

	return true
}

func check(e error) { if e != nil { panic(e) } }
