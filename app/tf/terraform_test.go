package main

import (
	"os"
	"testing"

	"github.com/battenworks/go-common/assert"
)

type FakeExecutor struct{}

func (fe FakeExecutor) Execute(cmdName string, cmdArgs ...string) ([]byte, error) {
	output := full_output

	return []byte(output), nil
}

func TestValidateWorkingDirectory(t *testing.T) {
	t.Run("succeeds when directory is valid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Create(backendFile)
		defer os.Remove(backendFile)

		actual, err := validateWorkingDirectory(currentDir)
		expected := currentDir

		assert.NoError(t, err)
		assert.Equals(t, expected, actual)
	})

	t.Run("fails when directory is invalid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Remove(backendFile)

		_, err := validateWorkingDirectory(currentDir)

		assert.True(t, err == ErrInvalidWorkingDirectory, "expected error '%s', received none", ErrInvalidWorkingDirectory)
	})
}

const full_output = `Note: Objects have changed outside of Terraform

\x1b[0mTerraform detected the following changes made outside of Terraform since the
last "terraform apply":

...
truncated
...

\x1b[0mUnless you have made equivalent changes to your configuration, or ignored the
\x1b[0mrelevant attributes using ignore_changes, the following plan may include
\x1b[0mactions to undo or respond to these changes.

─────────────────────────────────────────────────────────────────────────────

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan.`

const output_with_drift_removed = `Note: Objects have changed outside of Terraform

---- 12 lines hidden ----

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan.`

func TestPlan(t *testing.T) {
	t.Run("does NOT remove drift output when hideDrift is false", func(t *testing.T) {
		executor := FakeExecutor{}
		hideDrift := false

		actual, _ := plan(executor, hideDrift)
		expected := full_output

		assert.Equals(t, expected, actual)
	})

	t.Run("removes drift output when hideDrift is true", func(t *testing.T) {
		executor := FakeExecutor{}
		hideDrift := true

		actual, _ := plan(executor, hideDrift)
		expected := output_with_drift_removed

		assert.Equals(t, expected, actual)
	})
}
