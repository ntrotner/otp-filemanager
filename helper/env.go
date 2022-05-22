package helper

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getPort() *string {
	port := os.Getenv("HTTPPORT")

	if port == "" {
		port = "8080"
	}
	return &port
}

func getIssuer() *string {
	issuer := os.Getenv("ISSUER")
	if issuer == "" {
		issuer = "OTP-File-Manager"
	}
	return &issuer
}

func getPeriod() *uint {
	period, err := strconv.ParseUint(os.Getenv("PERIOD"), 10, 64)

	if err != nil {
		u_period := uint(30)
		return &u_period
	} else {
		u_period := uint(period)
		return &u_period
	}
}

func getIDSeed() *uint64 {
	id_seed, err := strconv.ParseUint(os.Getenv("IDSEED"), 10, 64)

	if err != nil {
		id_seed = uint64(42)
	}
	return &id_seed
}

func ReadEnv() (*string, *uint64, *string, *uint) {
	godotenv.Load(".env")

	port := getPort()
	id_seed := getIDSeed()
	issuer := getIssuer()
	period := getPeriod()

	return port, id_seed, issuer, period
}
