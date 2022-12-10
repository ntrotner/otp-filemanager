package common

import (
	"mime/multipart"
	"net/http"

	"github.com/kennygrant/sanitize"
)

func SanitizeFileName(fileName *string) {
	tempName := sanitize.Name(*fileName)
	tempName = sanitize.Path(tempName)

	*fileName = tempName
}

func SanitizeUploadFile(request *http.Request) (string, string, multipart.File, *multipart.FileHeader, error) {
	request.ParseMultipartForm(10 << 25)
	file, header, err := request.FormFile("file")

	return request.FormValue("user"), request.FormValue("password"), file, header, err
}
