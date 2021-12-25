package main

import (
	"command"
	"console"
	"errors"
	"fmt"
	"os"
)

func cleanScope(scope string, workspace string) error {
	workingDir, err := validateWorkingDirectory(scope)
	if err != nil {
		return err
	}

	fmt.Print("directory: ")
	console.White(workingDir + "\n")
	fmt.Print("workspace: ")
	console.White(workspace + "\n")

	fmt.Println("removing .terraform directory")
	err = cleanTerraformCacheDirectory(workingDir)
	if err != nil {
		return err
	}
	console.Green(".terraform directory removed\n")

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
	if _, err := os.Stat(dir + "/backend.tf"); os.IsNotExist(err) {
		return dir, ErrInvalidWorkingDirectory
	}

	return dir, nil
}

func cleanTerraformCacheDirectory(dir string) error {
	err := os.RemoveAll(dir + "/.terraform")
	if err != nil {
		return err
	}

	return nil
}

func initializeTerraform() ([]byte, error) {
	cmdName := "terraform"
	cmdArgs := []string{"init"}

	return command.Execute(cmdName, cmdArgs)
}

func selectWorkspace(workspace string) ([]byte, error) {
	cmdName := "terraform"
	cmdArgs := []string{"workspace", "select", workspace}

	return command.Execute(cmdName, cmdArgs)
}
