package http_api

import "net/http"

func MiscHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// send client that service is running
		// could be further extended to check if otp is available
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("http-api/html-templates/assets/"))))
}
