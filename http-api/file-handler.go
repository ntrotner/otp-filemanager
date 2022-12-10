package http_api

import (
	"log"
	"net/http"
	"otp-filemanager/http-api/common"
)

// FileHandler initialize file related endpoints
func FileHandler() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// check if user exists and code is valid
		err := common.ChallengeWriteFileHTTP(r, w)
		if err != nil {
			return
		}

		w.WriteHeader(200)
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		// check if user exists and code is valid
		foundFile, err := common.ChallengeReadFileHTTP(r, w)
		if err != nil {
			return
		}

		log.Println("Send File")
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(*foundFile)
	})
}
