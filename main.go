package main

import (
	"os"
	http_api "otp-filemanager/http-api"

	"github.com/joho/godotenv"
)

func main() {
	port := readEnv()
	http_api.InitializeHTTPServer(port)
}

func readEnv() *string {
	godotenv.Load(".env")

	port := os.Getenv("HTTPPORT")
	return &port
}
