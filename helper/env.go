package helper

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getPort() *string {
	port := os.Getenv("HTTPPORT")

	if port == "" {
		// set default if nothing is set
		port = "8080"
	}
	return &port
}

func getIssuer() *string {
	issuer := os.Getenv("ISSUER")
	if issuer == "" {
		// set default name if nothing is set
		issuer = "OTP-File-Manager"
	}
	return &issuer
}

func getPeriod() *uint {
	period, err := strconv.ParseUint(os.Getenv("PERIOD"), 10, 32)
	var u_period uint

	if err != nil {
		// set length of a valid password to 30 seconds by default
		u_period = uint(30)
	} else {
		u_period = uint(period)
	}
	return &u_period
}

func getIDSeed() *uint64 {
	// parse seed to uint64 for generating identities
	id_seed, err := strconv.ParseUint(os.Getenv("IDSEED"), 10, 64)

	if err != nil {
		id_seed = uint64(42)
	}
	return &id_seed
}

// get required env variables for running the service
func ReadEnv() (*string, *uint64, *string, *uint) {
	godotenv.Load(".env")

	port := getPort()
	id_seed := getIDSeed()
	issuer := getIssuer()
	period := getPeriod()

	return port, id_seed, issuer, period
}
