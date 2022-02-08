package main

import (
	"errors"
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
	t.Run("removes drift output", func(t *testing.T) {
		executor := FakeExecutor{}

		actual := quietPlan(executor)
		expected := output_with_drift_removed

		assert.Equals(t, expected, actual)
	})
}

func TestCanTurnFileOff(t *testing.T) {
	t.Run("returns false for backend file", func(t *testing.T) {
		file := "backend.tf"
		assert.False(t, canTurnFileOff(file), "should NOT be able to turn %s off", file)
	})
	t.Run("returns false for lock file", func(t *testing.T) {
		file := ".terraform.locl.hcl"
		assert.False(t, canTurnFileOff(file), "should NOT be able to turn %s off", file)
	})
	t.Run("returns true for files that have the TF extension", func(t *testing.T) {
		file1 := "file1.tf"
		file2 := "file2.tf"
		assert.True(t, canTurnFileOff(file1), "should be able to turn %s off", file1)
		assert.True(t, canTurnFileOff(file2), "should be able to turn %s off", file2)
	})
	t.Run("returns false for files that DONT have the TF extension", func(t *testing.T) {
		file1 := "foo.bar"
		file2 := "bar.baz"
		assert.False(t, canTurnFileOff(file1), "should NOT be able to turn %s off", file1)
		assert.False(t, canTurnFileOff(file2), "should NOT be able to turn %s off", file2)
	})
}

func TestCanTurnFileOn(t *testing.T) {
	t.Run("returns true for files that have the OFF extension", func(t *testing.T) {
		file1 := "file1.tf" + OffFileExtension
		file2 := "file2.tf" + OffFileExtension
		assert.True(t, canTurnFileOn(file1), "should be able to turn %s on", file1)
		assert.True(t, canTurnFileOn(file2), "should be able to turn %s on", file2)
	})
	t.Run("returns false for files that DONT have the OFF extension", func(t *testing.T) {
		backendFile := "backend.tf"
		lockFile := ".terraform.lock.hcl"
		assert.False(t, canTurnFileOn(backendFile), "should NOT be able to turn %s on", backendFile)
		assert.False(t, canTurnFileOn(lockFile), "should NOT be able to turn %s on", lockFile)
	})
}

func assertFileExists(tb testing.TB, file string) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		tb.FailNow()
	}
}

func assertFileNotExists(tb testing.TB, file string) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		// pass
	} else {
		tb.FailNow()
	}
}

func TestOff(t *testing.T) {
	t.Run("ignore backend file", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Create(backendFile)
		defer os.Remove(backendFile)

		err := off(currentDir)

		assert.NoError(t, err)
		assertFileExists(t, backendFile)
	})
	t.Run("ignore lock file", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		lockFile := currentDir + "/.terraform.lock.hcl"
		os.Create(lockFile)
		defer os.Remove(lockFile)

		err := off(currentDir)

		assert.NoError(t, err)
		assertFileExists(t, lockFile)
	})
	t.Run("adds OFF extension to TF files", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		file1 := currentDir + "/one.tf"
		file1off := file1 + OffFileExtension
		file2 := currentDir + "/two.tf"
		file2off := file2 + OffFileExtension
		os.Create(file1)
		os.Create(file2)
		defer os.Remove(file1)
		defer os.Remove(file2)

		err := off(currentDir)
		defer os.Remove(file1off)
		defer os.Remove(file2off)

		assert.NoError(t, err)
		assertFileNotExists(t, file1)
		assertFileNotExists(t, file2)
		assertFileExists(t, file1off)
		assertFileExists(t, file2off)
	})
}

func TestOn(t *testing.T) {
	t.Run("removes OFF extension from TF files", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		file1 := currentDir + "/one.tf"
		file1off := file1 + OffFileExtension
		file2 := currentDir + "/two.tf"
		file2off := file2 + OffFileExtension
		os.Create(file1off)
		os.Create(file2off)
		defer os.Remove(file1off)
		defer os.Remove(file2off)

		err := on(currentDir)
		defer os.Remove(file1)
		defer os.Remove(file2)

		assert.NoError(t, err)
		assertFileNotExists(t, file1off)
		assertFileNotExists(t, file2off)
		assertFileExists(t, file1)
		assertFileExists(t, file2)
	})
}
