package main

import (
	"command"
	"console"
	"errors"
	"fmt"
	"os"
)

var cmdName string = "terraform"

func passThrough(cmdArgs []string) error {
	result, err := command.Execute(cmdName, cmdArgs)
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func cleanScope(scope string, workspace string) error {
	workingDir, err := validateWorkingDirectory(scope)
	if err != nil {
		return err
	}

	fmt.Print("directory: ")
	console.White(workingDir + "\n")
	fmt.Print("workspace: ")
	console.White(workspace + "\n")

	fmt.Println("removing terraform cache")
	err = cleanTerraformCache(workingDir)
	if err != nil {
		return err
	}
	console.Green("terraform cache removed\n")

	fmt.Println("initializing terraform")
	result, err := initializeTerraform()
	if err != nil {
		fmt.Print(string(result))
		return err
	}
	console.Green("terraform initialized\n")

	fmt.Println("selecting workspace:", workspace)
	result, err = selectWorkspace(workspace)
	if err != nil {
		fmt.Print(string(result))
		return err
	}
	fmt.Print(string(result))

	return nil
}

// ErrInvalidWorkingDirectory error constant
var ErrInvalidWorkingDirectory = errors.New("invalid working directory: no backend.tf found")

func validateWorkingDirectory(dir string) (string, error) {
	if _, err := os.Stat(dir + "/backend.tf"); errors.Is(err, os.ErrNotExist) {
		return dir, ErrInvalidWorkingDirectory
	}

	return dir, nil
}

func cleanTerraformCache(dir string) error {
	err := os.RemoveAll(dir + "/.terraform")
	if err != nil {
		return err
	}

	if _, err := os.Stat(dir + "/.terraform.lock.hcl"); errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		err = os.Remove(dir + "/.terraform.lock.hcl")
		if err != nil {
			return err
		}
	}

	return nil
}

func initializeTerraform() ([]byte, error) {
	cmdArgs := []string{"init"}

	return command.Execute(cmdName, cmdArgs)
}

func selectWorkspace(workspace string) ([]byte, error) {
	cmdArgs := []string{"workspace", "select", workspace}

	return command.Execute(cmdName, cmdArgs)
}
