package file_system

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"otp-filemanager/entity-controller/security"
	"path"
)

func WriteFile(id *string, fileName *string, file *multipart.File) error {
	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return err
	}

	openedFile, err := os.Create(path.Join(PathToFilesOfIdentities, *id, *fileName))
	if err != nil {
		return err
	}

	fileBytes_ptr, err := security.Encrypt(&fileBytes)
	if err != nil {
		return err
	}

	_, err = openedFile.Write(*fileBytes_ptr)

	return err
}

func ReadFile(id *string, name *string) (*[]byte, error) {
	fo, err := os.ReadFile(path.Join(PathToFilesOfIdentities, *id, *name))
	if err != nil {
		return nil, err
	}
	return security.Decrypt(&fo)
}

func DeleteFile(id *string, name *string) error {
	return errors.New("not implemented")
}
