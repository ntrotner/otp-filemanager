package main

import (
	"otp-filemanager/helper"
	httpapi "otp-filemanager/http-api"
	permissioncontroller "otp-filemanager/permission-controller"
)

func main() {
	port, idSeed, issuer, period := helper.ReadEnv()

	// initialize modules
	permissioncontroller.InitializeOTPGenerator(idSeed, issuer, period)
	httpapi.InitializeHTTPServer(port)
}
