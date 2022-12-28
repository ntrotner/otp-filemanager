package main

import (
	"otp-filemanager/helper"
	httpapi "otp-filemanager/http-api"
	permissioncontroller "otp-filemanager/permission-controller"
)

func main() {
	port, idSeed, issuer, period, maxSize, modifier := helper.ReadEnv()

	// initialize modules
	permissioncontroller.InitializeOTPGenerator(modifier, idSeed, issuer, period)
	httpapi.InitializeHTTPServer(port, maxSize)
}
