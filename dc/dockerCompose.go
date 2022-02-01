package main

import (
	"fmt"
)

const cmdName = "docker-compose"

func passThrough(executor Executor, cmdArgs []string) error {
	result, err := executor.Execute(cmdName, cmdArgs)
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}

func upDisconnected(executor Executor) error {
	cmdArgs := []string{"up", "-d"}
	result, err := executor.Execute(cmdName, cmdArgs)
	fmt.Print(string(result))

	if err != nil {
		return err
	}

	return nil
}
