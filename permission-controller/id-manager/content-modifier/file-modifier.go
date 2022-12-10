package content_modifier

import (
	"io"
	"mime/multipart"
	"os"
)

func WriteFile(fileName string, file *multipart.File) error {
	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return err
	}

	openedFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = openedFile.Write(fileBytes)

	return err
}

func ReadFile(name string) (*[]byte, error) {
	fo, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &fo, nil
}

func DeleteFile() {
}
