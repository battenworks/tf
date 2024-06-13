package main

import (
	"os"
	"strings"

	"github.com/battenworks/go-common/v2/console"
	"github.com/battenworks/tf/v2/tfcmd"
)

var version string = "built from source"

func main() {
	readable_version := "Version: " + strings.Replace(version, "v", "", -1)

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "version", "-v", "-version", "--version":
			console.Outln(readable_version)
		case "clean":
			workingDir := getWorkingDirectory()

			err := tfcmd.ValidateWorkingDirectory(workingDir)
			if err != nil {
				console.Outln(err.Error())
				os.Exit(1)
			}

			console.Outln("removing terraform cache")
			err = tfcmd.CleanTerraformCache(workingDir)
			if err != nil {
				os.Exit(1)
			}

			console.Greenln("terraform cache removed")
			console.Outln("initializing terraform")
			err = tfcmd.PassThrough([]string{"init"})
			if err != nil {
				os.Exit(1)
			}
		case "wipe":
			workingDir := getWorkingDirectory()

			console.Outln("removing terraform cache")
			err := tfcmd.CleanTerraformCache(workingDir)
			if err != nil {
				os.Exit(1)
			}

			console.Greenln("terraform cache removed")
		case "off":
			workingDir := getWorkingDirectory()

			err := tfcmd.ValidateWorkingDirectory(workingDir)
			if err != nil {
				console.Outln(err.Error())
				os.Exit(1)
			}

			err = tfcmd.Off(workingDir)
			if err != nil {
				os.Exit(1)
			}
		case "on":
			workingDir := getWorkingDirectory()

			err := tfcmd.ValidateWorkingDirectory(workingDir)
			if err != nil {
				console.Outln(err.Error())
				os.Exit(1)
			}

			err = tfcmd.On(workingDir)
			if err != nil {
				os.Exit(1)
			}
		case "test":
			console.Outln("fmt > init > upgrade")
			err := tfcmd.PassThrough([]string{"fmt", "-recursive"})
			if err != nil {
				os.Exit(1)
			}

			err = tfcmd.PassThrough([]string{"init", "-upgrade"})
			if err != nil {
				os.Exit(1)
			}
			
			err = tfcmd.PassThrough([]string{"test"})
			if err != nil {
				os.Exit(1)
			}
		case "help", "-help", "--help":
			usage(readable_version)
		default:
			err := tfcmd.PassThrough(os.Args[1:])
			if err != nil {
				os.Exit(1)
			}
		}
	} else {
		usage(readable_version)
	}
}

func getWorkingDirectory() string {
	workingDir, err := os.Getwd()
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
	}

	return workingDir
}

func usage(readable_version string) {
	console.Whiteln("Wrapper for the Terraform CLI")
	console.Whiteln("Provides some opinionated commands to help with Terraform CLI use")
	console.Whiteln("All other commands are passed directly to the Terraform CLI")
	console.Outln("")
	console.Whiteln(readable_version)
	console.Outln("")
	console.Whiteln("Usage: tf COMMAND")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("clean")
	console.Whiteln("\t- Removes the Terraform cache and lock file from the current scope, then runs 'init'")
	console.Yellow("test")
	console.Whiteln("\t- Runs 'fmt -recursive', then 'init -upgrade', then 'test'")
	console.Yellow("wipe")
	console.Whiteln("\t- Removes the Terraform cache and lock file from the current scope")
	console.Yellow("off")
	console.Whiteln("\t- Adds the '.off' extension to all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to destroy all resources in the current scope")
	console.Yellow("on")
	console.Whiteln("\t- Removes the '.off' extension from all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to re-create all resources in the current scope")
}
