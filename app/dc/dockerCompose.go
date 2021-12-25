package main

import (
	"command"
	"fmt"
)

var cmdName string = "docker-compose"

func passThrough(cmdArgs []string) error {
	result, err := command.Execute(cmdName, cmdArgs)
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func upDisconnected() error {
	cmdArgs := []string{"up", "-d"}

	result, err := command.Execute(cmdName, cmdArgs)
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func fixAuthorization() error {
	result, err := stopHealer()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	result, err = stopListener()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	result, err = startHealer()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	result, err = startListener()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func stopHealer() ([]byte, error) {
	cmdArgs := []string{"stop", "healer"}

	return command.Execute(cmdName, cmdArgs)
}

func startHealer() ([]byte, error) {
	cmdArgs := []string{"start", "healer"}

	return command.Execute(cmdName, cmdArgs)
}

func stopListener() ([]byte, error) {
	cmdArgs := []string{"stop", "listener"}

	return command.Execute(cmdName, cmdArgs)
}

func startListener() ([]byte, error) {
	cmdArgs := []string{"start", "listener"}

	return command.Execute(cmdName, cmdArgs)
}
