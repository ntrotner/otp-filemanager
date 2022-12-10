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

func SanitizeUploadFile(request *http.Request, maxSize *int64) (string, string, multipart.File, *multipart.FileHeader, error) {
	request.ParseMultipartForm((2 << 19) * *maxSize)
	file, header, err := request.FormFile("file")

	return request.FormValue("user"), request.FormValue("password"), file, header, err
}
