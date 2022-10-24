package helper

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// CreateDirectory creates new folder at the root directory of the executable
func CreateDirectory(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return os.MkdirAll(folder, 0755)
	}
	return nil
}

// FilenameWithoutExtension removes the extension of the filename
func FilenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}

// ReadFileNamesOfDirectory returns list of files in a directory without child directories
func ReadFileNamesOfDirectory(folder *string) []string {
	fileNames := make([]string, 0)

	files, err := ioutil.ReadDir(*folder)
	if err != nil {
		return fileNames
	}

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}
