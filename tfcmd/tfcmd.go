package tfcmd

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const cmdName = "tofu"

// ErrInvalidWorkingDirectory error constant.
var ErrInvalidWorkingDirectory = errors.New("invalid working directory: no providers.tf found")

// OffFileExtension is the file extension used to turn .tf files off and on.
var OffFileExtension = ".off"

// ValidateWorkingDirectory determines if the supplied directory can be manipulated by this app.
func ValidateWorkingDirectory(dir string) error {
	if _, err := os.Stat(dir + "/providers.tf"); errors.Is(err, os.ErrNotExist) {
		return ErrInvalidWorkingDirectory
	}

	return nil
}

// CleanCache removes module cache and lock files from the supplied directory.
func CleanCache(dir string) error {
	err := os.RemoveAll(dir + "/.terraform")
	if err != nil {
		return err
	}

	if _, err := os.Stat(dir + "/.terraform.lock.hcl"); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	err = os.Remove(dir + "/.terraform.lock.hcl")
	if err != nil {
		return err
	}

	return nil
}

// Off adds a file extension to select config files,
// effectively turning them off for subsequent CLI operations.
func Off(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && CanTurnFileOff(file.Name()) {
			filePath := dir + "/" + file.Name()

			err := os.Rename(filePath, filePath+OffFileExtension)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CanTurnFileOff determines whether or not a file can have the "off" file extension.
func CanTurnFileOff(file string) bool {
	return file != "backend.tf" &&
		file != "backend_override.tf" &&
		file != "providers.tf" &&
		file != "variables.tf" &&
		filepath.Ext(file) == ".tf"
}

// On removes the file extensions that makes the CLI ignore config files,
// effectively turning them on for subsequent CLI operations.
func On(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && CanTurnFileOn(file.Name()) {
			filePath := dir + "/" + file.Name()

			err := os.Rename(filePath, filePath[0:len(filePath)-len(OffFileExtension)])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CanTurnFileOn determines whether or not a file can have the "off" extenstion removed.
func CanTurnFileOn(file string) bool {
	return filepath.Ext(file) == OffFileExtension
}

// PassThrough simply passes the commands to the CLI, unmodified.
func PassThrough(cmdArgs []string) error {
	cmd := exec.Command(cmdName, cmdArgs...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	return cmd.Run()
}
