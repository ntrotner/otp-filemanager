package common

import (
	"log"
	"net/http"
	entity_controller "otp-filemanager/entity-controller"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
	"time"
)

func extractCredentials(request *http.Request) (*string, *string) {
	request.ParseForm()
	id := request.FormValue("user")
	clientOverlappingCode := request.FormValue("password")

	return &id, &clientOverlappingCode
}

func extractFileName(request *http.Request) *string {
	request.ParseForm()
	fileName := request.FormValue("fileName")
	SanitizeFileName(&fileName)

	return &fileName
}

func ChallengeLoginHTTP(request *http.Request, writer http.ResponseWriter) (*content_modifier.UserOtp, error) {
	// get username and password
	id, clientOverlappingCode := extractCredentials(request)

	currentTime := time.Now()

	// check if user exists and code is valid
	foundID, err := entity_controller.ChallengeLogin(id, clientOverlappingCode, &currentTime)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("Login Failed", *id)
		return nil, err
	}

	return foundID, nil
}

func ChallengeReadFileHTTP(request *http.Request, writer http.ResponseWriter) (*[]byte, error) {
	// get username and password
	id, clientOverlappingCode := extractCredentials(request)
	fileName := extractFileName(request)

	currentTime := time.Now()

	// check if user exists and code is valid
	foundFile, err := entity_controller.ChallengeReadFile(id, clientOverlappingCode, &currentTime, fileName)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("Read File Failed", *id)
		return nil, err
	}

	return foundFile, nil
}

func ChallengeWriteFileHTTP(request *http.Request, writer http.ResponseWriter, maxSize *int64) error {
	// get username and password
	id, clientOverlappingCode, file, handler, err := SanitizeUploadFile(request, maxSize)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("Write File Failed for Uploaded File", err)
		return err
	}

	currentTime := time.Now()

	// check if user exists and code is valid
	err = entity_controller.ChallengeWriteFile(&id, &clientOverlappingCode, &currentTime, &handler.Filename, &file)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("Write File Failed", id, "=", err)
		return err
	}

	return nil
}

func ChallengeDeleteIdentity(request *http.Request, writer http.ResponseWriter) (*string, error) {
	// get username and password
	id, clientOverlappingCode := extractCredentials(request)

	currentTime := time.Now()

	// check if user exists and code is valid, then delete all related files and user
	err := entity_controller.ChallengeDeleteIdentity(id, clientOverlappingCode, &currentTime)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("Identity Deletion Failed", *id)
		return id, err
	}

	return id, nil
}

func ChallengeDeleteFile(request *http.Request, writer http.ResponseWriter) error {
	// get username and password
	id, clientOverlappingCode := extractCredentials(request)
	fileName := extractFileName(request)

	currentTime := time.Now()

	// check if user exists and code is valid, then delete selected file
	err := entity_controller.ChallengeDeleteFile(id, clientOverlappingCode, &currentTime, fileName)
	if err != nil {
		writer.WriteHeader(401)
		writer.Write([]byte("Access Denied"))
		log.Println("File Deletion Failed", *id)
		return err
	}

	return nil
}
