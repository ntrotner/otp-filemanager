package main

import (
	entity_controller "otp-filemanager/entity-controller"
	"otp-filemanager/helper"
	http_api "otp-filemanager/http-api"
)

func main() {
	port, idSeed, issuer, period, maxSize, modifier, key, expirationTime := helper.ReadEnv()

	// initialize modules
	entity_controller.InitializeOTPGenerator(key, modifier, idSeed, issuer, period, expirationTime)
	http_api.InitializeHTTPServer(port, maxSize)
}
