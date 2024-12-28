package tfcmd_test

import (
	"errors"
	"os"
	"testing"

	"github.com/battenworks/go-common/v3/assert"
	"github.com/battenworks/tf/v2/tfcmd"
)

func TestValidateWorkingDirectory(t *testing.T) {
	t.Run("succeeds when directory is valid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Create(backendFile)
		defer os.Remove(backendFile)

		err := tfcmd.ValidateWorkingDirectory(currentDir)

		assert.NoError(t, err)
	})

	t.Run("fails when directory is invalid", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		backendFile := currentDir + "/backend.tf"
		os.Remove(backendFile)

		err := tfcmd.ValidateWorkingDirectory(currentDir)

		assert.True(t, err == tfcmd.ErrInvalidWorkingDirectory, "expected error '%s', received none", tfcmd.ErrInvalidWorkingDirectory)
	})
}

func TestCanTurnFileOff(t *testing.T) {
	t.Run("returns false for backend file", func(t *testing.T) {
		file := "backend.tf"
		assert.False(t, tfcmd.CanTurnFileOff(file), "should NOT be able to turn %s off", file)
	})
	t.Run("returns false for providers file", func(t *testing.T) {
		file := "providers.tf"
		assert.False(t, tfcmd.CanTurnFileOff(file), "should NOT be able to turn %s off", file)
	})
	t.Run("returns false for lock file", func(t *testing.T) {
		file := ".terraform.locl.hcl"
		assert.False(t, tfcmd.CanTurnFileOff(file), "should NOT be able to turn %s off", file)
	})
	t.Run("returns true for files that have the TF extension", func(t *testing.T) {
		file1 := "file1.tf"
		file2 := "file2.tf"
		assert.True(t, tfcmd.CanTurnFileOff(file1), "should be able to turn %s off", file1)
		assert.True(t, tfcmd.CanTurnFileOff(file2), "should be able to turn %s off", file2)
	})
	t.Run("returns false for files that DONT have the TF extension", func(t *testing.T) {
		file1 := "foo.bar"
		file2 := "bar.baz"
		assert.False(t, tfcmd.CanTurnFileOff(file1), "should NOT be able to turn %s off", file1)
		assert.False(t, tfcmd.CanTurnFileOff(file2), "should NOT be able to turn %s off", file2)
	})
}

func TestCanTurnFileOn(t *testing.T) {
	t.Run("returns true for files that have the OFF extension", func(t *testing.T) {
		file1 := "file1.tf" + tfcmd.OffFileExtension
		file2 := "file2.tf" + tfcmd.OffFileExtension
		assert.True(t, tfcmd.CanTurnFileOn(file1), "should be able to turn %s on", file1)
		assert.True(t, tfcmd.CanTurnFileOn(file2), "should be able to turn %s on", file2)
	})
	t.Run("returns false for files that DONT have the OFF extension", func(t *testing.T) {
		backendFile := "backend.tf"
		lockFile := ".terraform.lock.hcl"
		assert.False(t, tfcmd.CanTurnFileOn(backendFile), "should NOT be able to turn %s on", backendFile)
		assert.False(t, tfcmd.CanTurnFileOn(lockFile), "should NOT be able to turn %s on", lockFile)
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

		err := tfcmd.Off(currentDir)

		assert.NoError(t, err)
		assertFileExists(t, backendFile)
	})
	t.Run("ignore lock file", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		lockFile := currentDir + "/.terraform.lock.hcl"
		os.Create(lockFile)
		defer os.Remove(lockFile)

		err := tfcmd.Off(currentDir)

		assert.NoError(t, err)
		assertFileExists(t, lockFile)
	})
	t.Run("adds OFF extension to TF files", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		file1 := currentDir + "/one.tf"
		file1off := file1 + tfcmd.OffFileExtension
		file2 := currentDir + "/two.tf"
		file2off := file2 + tfcmd.OffFileExtension
		os.Create(file1)
		os.Create(file2)
		defer os.Remove(file1)
		defer os.Remove(file2)

		err := tfcmd.Off(currentDir)
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
		file1off := file1 + tfcmd.OffFileExtension
		file2 := currentDir + "/two.tf"
		file2off := file2 + tfcmd.OffFileExtension
		os.Create(file1off)
		os.Create(file2off)
		defer os.Remove(file1off)
		defer os.Remove(file2off)

		err := tfcmd.On(currentDir)
		defer os.Remove(file1)
		defer os.Remove(file2)

		assert.NoError(t, err)
		assertFileNotExists(t, file1off)
		assertFileNotExists(t, file2off)
		assertFileExists(t, file1)
		assertFileExists(t, file2)
	})
}
