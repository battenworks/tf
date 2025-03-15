package main

import (
	"os"
	"strings"

	"github.com/battenworks/go-common/v3/console"
	"github.com/battenworks/tf/v2/tfcmd"
)

var version string = "built from source"

func main() {
	readable_version := "Version: " + strings.Replace(version, "v", "", -1)

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "clean":
			workingDir := getWorkingDirectory()

			err := tfcmd.ValidateWorkingDirectory(workingDir)
			if err != nil {
				console.Outln(err.Error())
				os.Exit(1)
			}

			console.Outln("removing tofu cache")
			err = tfcmd.CleanCache(workingDir)
			if err != nil {
				os.Exit(1)
			}

			console.Greenln("tofu cache removed")
			console.Outln("initializing tofu")
			err = tfcmd.PassThrough([]string{"init"})
			if err != nil {
				os.Exit(1)
			}
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
		case "replan":
			workingDir := getWorkingDirectory()

			err := tfcmd.ValidateWorkingDirectory(workingDir)
			if err != nil {
				console.Outln(err.Error())
				os.Exit(1)
			}

			err = tfcmd.PassThrough([]string{"init", "-upgrade"})
			if err != nil {
				os.Exit(1)
			}

			err = tfcmd.PassThrough([]string{"plan"})
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
		case "version", "-v", "-version", "--version":
			console.Outln(readable_version)
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
	console.Whiteln("Wrapper for the OpenTofu CLI")
	console.Whiteln("Provides some opinionated commands to help with CLI use")
	console.Whiteln("All other commands are passed directly to the CLI")
	console.Outln("")
	console.Whiteln(readable_version)
	console.Outln("")
	console.Whiteln("Usage: tf COMMAND")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("clean")
	console.Whiteln("\t- Removes the cache and lock file from the current scope, then runs 'init'")
	console.Yellow("off")
	console.Whiteln("\t- Adds the '.off' extension to all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to destroy all resources in the current scope")
	console.Yellow("on")
	console.Whiteln("\t- Removes the '.off' extension from all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to re-create all resources in the current scope")
	console.Yellow("replan")
	console.Whiteln("\t- Runs 'init -upgrade', then 'plan'")
	console.Yellow("test")
	console.Whiteln("\t- Runs 'fmt -recursive', then 'init -upgrade', then 'test'")
}
