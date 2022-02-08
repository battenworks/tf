package main

import (
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

		switch command {
		case "clean":
			workingDir:= getWorkingDirectory()

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
		case "qplan":
			result := quietPlan(executor)

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
		os.Exit(1)
	}

	workingDir, err := validateWorkingDirectory(scope)
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
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
	console.Yellow("  qplan")
	console.Whiteln(" - calls terraform plan and hides drift output that results from the refresh stage of the plan")
	console.Yellow("  off")
	console.Whiteln(" - adds the '.off' extension to all config files in the working directory")
	console.Yellow("  on")
	console.Whiteln(" - removes the '.off' extension from all config files in the working directory")
}
