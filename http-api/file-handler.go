package http_api

import "net/http"

// FileHandler initialize file related endpoints
func FileHandler() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {

	})
}
