package permission_controller

import (
	"log"
	"otp-filemanager/helper"
	idmanager "otp-filemanager/permission-controller/id-manager"
	contentmodifier "otp-filemanager/permission-controller/id-manager/content-modifier"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/teris-io/shortid"
)

var (
	idGenerator     *shortid.Shortid
	GenerateOtpOpts totp.GenerateOpts
	ValidateOtpOpts totp.ValidateOpts
)

const SAFE_CHARACTERS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

// InitializeOTPGenerator create worker for generating identities and default options for otp
func InitializeOTPGenerator(modifier *int8, seed *uint64, issuer *string, period *uint) {
	idmanager.InitializeIDManager(modifier)
	idGenerator, _ = shortid.New(1, SAFE_CHARACTERS, *seed)

	GenerateOtpOpts = totp.GenerateOpts{
		Issuer:      *issuer,
		AccountName: "",
		Period:      *period,
		SecretSize:  0,
		Secret:      nil,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        nil,
	}

	ValidateOtpOpts = totp.ValidateOpts{
		Period:    *period,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	}
}

// create new random identity
func CreateIdentity() (*contentmodifier.UserOtp, error) {
	// create new identity
	newId, err := idGenerator.Generate()
	if err != nil {
		return nil, err
	}

	// modify id created by generator
	helper.MapUserID(&newId)

	// check if identity exist and fail if it does
	userOtp, err := idmanager.ExistsIdentity(&newId)
	if err == nil {
		return userOtp, err
	}

	// copy default options and adjust name of new account
	otpUserOpts := GenerateOtpOpts
	otpUserOpts.AccountName = newId

	// generate otp key
	key, _ := totp.Generate(otpUserOpts)

	newUser := contentmodifier.UserOtp{
		Id:  newId,
		Key: *key,
	}

	// save new identity
	err = idmanager.CreateIdentity(&newId, &newUser)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &newUser, nil
}
