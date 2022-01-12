package main

import (
	"command"
	"fmt"
)

const cmdName = "docker-compose"

func passThrough(cmdArgs []string) error {
	commander := command.ExecCommander{Name: cmdName, Args: cmdArgs}
	result, err := commander.Execute()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func upDisconnected() error {
	cmdArgs := []string{"up", "-d"}
	commander := command.ExecCommander{Name: cmdName, Args: cmdArgs}
	result, err := commander.Execute()
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}
