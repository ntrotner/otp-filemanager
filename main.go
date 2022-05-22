package main

import (
	"otp-filemanager/helper"
	http_api "otp-filemanager/http-api"
	permission_controller "otp-filemanager/permission-controller"
)

func main() {
	port, id_seed, issuer, period := helper.ReadEnv()
	permission_controller.InitializeOTPGenerator(id_seed, issuer, period)
	http_api.InitializeHTTPServer(port)
}
