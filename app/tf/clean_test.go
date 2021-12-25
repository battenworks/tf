package main

import (
	"os"
	"testing"
)

func TestValidateWorkingDirectory(t *testing.T) {
	checkForError := func(t *testing.T, err error, message string) {
		t.Helper()
		if err != nil {
			t.Errorf(message+": %s", err)
		}
	}

	t.Run("succeeds when directory is valid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Create(backendFile)
		defer os.Remove(backendFile)

		actual, err := validateWorkingDirectory(currentDir)
		expected := currentDir

		checkForError(t, err, "error validating working directory")
		if actual != expected {
			t.Errorf("expected directory: '%s', actual directory: '%s'", expected, actual)
		}
	})

	t.Run("fails when directory is invalid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Remove(backendFile)

		_, err := validateWorkingDirectory(currentDir)

		if err != ErrInvalidWorkingDirectory {
			t.Errorf("expected error '%s', received none", ErrInvalidWorkingDirectory)
		}
	})
}
