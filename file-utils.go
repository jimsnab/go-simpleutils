package simpleutils

import (
	"errors"
	"os"

	"github.com/djherbis/atime"
)

//FileExists returns true if path corresponds to a file, and false
//if it corresponds to nothing or to a directory.
func FileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), nil
}

//IsDirectory returns true if path corresponds to a directory, and false
//if it corresponds to nothing or to a file.
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return fileInfo.IsDir(), nil
}

//CopyFile reads file content into memory and writes it to a new file
func CopyFile(srcPath, destPath string) (int64, error) {
	failed := false

	src, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return 0, err
	}
	defer func() {
		dest.Close()
		if failed {
			os.Remove(destPath)
		}
	}()

	failed = true

	len, err := dest.ReadFrom(src)
	if err != nil {
		return 0, err
	}

	err = dest.Sync()
	if err != nil {
		return 0, err
	}

	fi, err := os.Stat(srcPath)
	if err != nil {
		return 0, err
	}

	err = os.Chmod(destPath, fi.Mode())
	if err != nil {
		return 0, err
	}

	atime, err := atime.Stat(srcPath)
	if err != nil {
		return 0, err
	}

	err = os.Chtimes(destPath, atime, fi.ModTime())
	if err != nil {
		return 0, err
	}

	failed = false
	return len, nil
}
