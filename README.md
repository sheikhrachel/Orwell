# Orwell

Orwell is an open source developer productivity CLI tool built to manage projects from creation through production, built with Go and Cobra

## Features

- Gitignores pulled from Github bootstrap examples

- Sets up git connections based on provided repository remote

- Git hooks configured with Pre-Commit

- Current environment language support: Python3

- Current License support: Apache, GNUGPL v3.0, MIT

## Prerequisites

Ensure you have Go

```zsh
$ go version # returns go version go1.14 darwin/amd64 or greater
```

## Get Started

```zsh
# Install Orwell on your machine
$ go install orwell

# Validate your Orwell installation
$ orwell --help
```

## Usage

```zsh
# Create a project with Coldstart
# Navigate to your target parent directory and run the following
$ orwell coldstart -l {language} -x {license} -n {project name} -g {github repo url}
```

## Contact Info

```Go
name  := "Rachel Sheikh"
email := "sheikhrachel97@gmail.com"
```