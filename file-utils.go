package simpleutils

import "os"

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
		return false, err
	}
	return fileInfo.IsDir(), nil
}

