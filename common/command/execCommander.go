package command

import (
	"console"
	"os/exec"
	"strings"
)

type ExecCommander struct {
	Name string
	Args []string
}

// Execute executes the configured command with the configured arguments
func (ec ExecCommander) Execute() ([]byte, error) {
	cmd := exec.Command(ec.Name, ec.Args...)
	out, err := cmd.CombinedOutput()

	console.Yellow(strings.Join(cmd.Args, " "))

	if err != nil {
		console.Red(" error\n")
		return out, err
	}

	console.Green(" done\n")

	return out, nil
}
