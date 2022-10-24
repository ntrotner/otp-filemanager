package http_api

import (
	"log"
	"net/http"
)

// start http server to handle requests that relate to backend actions
func InitializeHTTPServer(port *string) {
	// initialize otp related endpoints
	OTPHandler()

	// import endpoints not directly related to otp
	MiscHandler()

	log.Println("Listening on Port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
