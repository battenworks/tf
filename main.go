package main

import (
	"flag"
	"os"

	"github.com/battenworks/go-common/command"
	"github.com/battenworks/go-common/console"
)

type Executor interface {
	Execute(cmdName string, cmdArgs ...string) ([]byte, error)
}

var executor Executor

func main() {
	executor = command.CommandExecutor{}

	if len(os.Args) > 1 {
		command := os.Args[1]

		cleanFlags := flag.NewFlagSet("clean", flag.ExitOnError)
		workspace := cleanFlags.String("workspace", "default", "Set the Terraform workspace to init and use")

		planFlags := flag.NewFlagSet("plan", flag.ExitOnError)
		hideDrift := planFlags.Bool("hide-drift", false, "Hide Terraform's verbose drift output")

		switch command {
		case "clean":
			cleanFlags.Parse(os.Args[2:])
			workingDir := getWorkingDirectory()

			console.Outln("removing terraform cache")
			err := cleanTerraformCache(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Greenln("terraform cache removed")

			console.Outln("initializing terraform")
			initResult, err := initializeTerraform(executor)
			console.Out(initResult)
			if err != nil {
				break
			}

			err = selectWorkspace(executor, *workspace)
			console.Outln("")
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out("workspace: ")
			console.WhitelnBold(*workspace)
		case "plan":
			planFlags.Parse(os.Args[2:])
			result := plan(executor, *hideDrift)
			console.Out(result)
		case "off":
			workingDir := getWorkingDirectory()

			err := off(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "on":
			workingDir := getWorkingDirectory()

			err := on(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "help", "-help", "--help":
			usage()
		default:
			result, _ := passThrough(executor, os.Args[1:])
			console.Out(result)
		}
	} else {
		usage()
	}
}

func getWorkingDirectory() string {
	scope, err := os.Getwd()
	if err != nil {
		console.Outln(err.Error())
	}

	workingDir, err := validateWorkingDirectory(scope)
	if err != nil {
		console.Outln(err.Error())
	}

	return workingDir
}

func usage() {
	console.Whiteln("Wrapper for the Terraform CLI")
	console.Whiteln("Provides some opinionated commands to help with Terraform CLI use")
	console.Whiteln("All other commands are passed directly to the Terraform CLI")
	console.Outln("")
	console.Whiteln("Usage: tf COMMAND [args]")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("  clean")
	console.Whiteln(" - wipes terraform cache from current scope, and re-inits")
	console.Whiteln("    args: -workspace=<workspace to select after initialization>")
	console.Yellow("  plan")
	console.Whiteln(" - calls terraform plan with an optional arg to hide drift output")
	console.Whiteln("    args: -hide-drift")
	console.Yellow("  off")
	console.Whiteln(" - adds the '.off' extension to all config files in the working directory")
	console.Yellow("  on")
	console.Whiteln(" - removes the '.off' extension from all config files in the working directory")
}
