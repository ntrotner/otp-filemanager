package helper

import (
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
