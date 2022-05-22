package http_api

import (
	"log"
	"net/http"
)

// start http server to handle requests that relate to backend actions
func InitializeHTTPServer(port *string) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// send client that service is running
		// could be further extended to check if otp is available
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"status": "ok"}`))
	})

	// initialize otp related endpoints
	OTPHandler()

	log.Println("Listening on Port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
