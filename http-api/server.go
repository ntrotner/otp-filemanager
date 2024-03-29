package http_api

import (
	"log"
	"net/http"
)

// start http server to handle requests that relate to backend actions
func InitializeHTTPServer(port *string, maxSize *int64) {
	// initialize otp related endpoints
	OTPHandler()

	// initialize endpoint to interact with files
	FileHandler(maxSize)

	// import endpoints not directly related to otp
	MiscHandler()

	log.Println("Listening on Port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
