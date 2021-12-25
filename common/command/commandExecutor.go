package command

import (
	"bytes"
	"console"
	"os/exec"
	"strings"
)

// Execute executes the given command with the given arguments
func Execute(commandName string, commandArgs []string) ([]byte, error) {
	cmd := exec.Command(commandName, commandArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	console.Yellow(strings.Join(cmd.Args, " "))

	err := cmd.Run()
	if err != nil {
		console.Red(" error\n")
		return stderr.Bytes(), err
	}

	console.Green(" done\n")

	return stdout.Bytes(), nil
}
