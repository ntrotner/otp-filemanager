package content_modifier

import (
	"mime/multipart"
)

type WriteFile func(id *string, fileName *string, file *multipart.File) error
type ReadFile func(id *string, name *string) (*[]byte, error)
type DeleteFile func(id *string, name *string) error

type FileModifier struct {
	WriteFile  WriteFile
	ReadFile   ReadFile
	DeleteFile DeleteFile
}
