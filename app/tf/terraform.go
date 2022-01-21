package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const cmdName = "terraform"

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

func initializeTerraform(executor Executor) (string, error) {
	cmdArgs := []string{"init"}
	result, err := executor.Execute(cmdName, cmdArgs...)

	return string(result), err
}

func selectWorkspace(executor Executor, workspace string) (string, error) {
	cmdArgs := []string{"workspace", "select", workspace}
	result, err := executor.Execute(cmdName, cmdArgs...)

	return string(result), err
}

func plan(executor Executor, hideDrift bool) (string, error) {
	cmdArgs := []string{"plan"}

	result, err := executor.Execute(cmdName, cmdArgs...)
	if err != nil {
		return string(result), err
	}

	if hideDrift {
		lines := strings.Split(string(result), "\n")
		linesDiscarded := 0
		filteredOutput := []string{}
		discarding := false
		regexBeginDiscard := regexp.MustCompile(`Terraform detected the following changes made outside of Terraform since the$`)
		regexEndDiscard1 := regexp.MustCompile(`Unless you have made equivalent changes to your configuration, or ignored the$`)
		regexEndDiscard2 := regexp.MustCompile(`relevant attributes using ignore_changes, the following plan may include$`)
		regexEndDiscard3 := regexp.MustCompile(`actions to undo or respond to these changes.$`)

		for i := 0; i < len(lines); i++ {

			if regexBeginDiscard.Match([]byte(lines[i])) {
				discarding = true
			}

			if discarding {
				linesDiscarded++

				if regexEndDiscard1.Match([]byte(lines[i])) &&
					regexEndDiscard2.Match([]byte(lines[i+1])) &&
					regexEndDiscard3.Match([]byte(lines[i+2])) {
					i = i + 4
					discarding = false
					filteredOutput = append(filteredOutput, "---- "+fmt.Sprint(linesDiscarded+4)+" lines hidden ----")
				}
			} else {
				filteredOutput = append(filteredOutput, lines[i])
			}
		}

		if !discarding {
			return strings.Join(filteredOutput, "\n"), nil
		}
	}

	return string(result), nil
}

func passThrough(executor Executor, cmdArgs []string) (string, error) {
	result, err := executor.Execute(cmdName, cmdArgs...)

	return string(result), err
}
