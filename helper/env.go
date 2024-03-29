package helper

import (
	"log"
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
	idSeed, err := strconv.ParseUint(os.Getenv("IDSEED"), 10, 64)

	if err != nil {
		idSeed = uint64(42)
	}
	return &idSeed
}

func getMaxSize() *int64 {
	maxSize, err := strconv.ParseInt(os.Getenv("MAXFILESIZEMB"), 10, 64)

	if err != nil {
		maxSize = int64(10)
	}
	return &maxSize
}

func getModifier() *int8 {
	modifier_i64, err := strconv.ParseInt(os.Getenv("MODIFIER"), 10, 8)
	modifier := int8(modifier_i64)

	if err != nil {
		modifier = int8(0)
	}
	return &modifier
}

func getKey() *string {
	key := os.Getenv("KEY")
	if key == "" {
		// set unsecure key
		key = ""
	}
	return &key
}

func getExpirationTime() *int64 {
	// parse minutes to int64
	expirationTime, err := strconv.ParseInt(os.Getenv("EXPIRATIONTIME"), 10, 64)

	if err != nil {
		expirationTime = int64(0)
	}
	return &expirationTime
}

// get required env variables for running the service
func ReadEnv() (*string, *uint64, *string, *uint, *int64, *int8, *string, *int64) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("reading .env file failed: ", err)
	}

	idSeed := getIDSeed()
	issuer := getIssuer()
	key := getKey()
	maxSize := getMaxSize()
	modifier := getModifier()
	period := getPeriod()
	port := getPort()
	expirationTime := getExpirationTime()

	return port, idSeed, issuer, period, maxSize, modifier, key, expirationTime
}
