package permission_controller

import (
	"log"
	id_manager "otp-filemanager/entity-controller/id-manager"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
	"otp-filemanager/entity-controller/security"
	"otp-filemanager/helper"

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
func InitializeOTPGenerator(key *string, modifier *int8, seed *uint64, issuer *string, period *uint) {
	security.InitializeSecurity(key)
	id_manager.InitializeIDManager(modifier)
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
func CreateIdentity() (*content_modifier.UserOtp, error) {
	// create new identity
	newId, err := idGenerator.Generate()
	if err != nil {
		return nil, err
	}

	// modify id created by generator
	helper.MapUserID(&newId)

	// check if identity exist and fail if it does
	userOtp, err := id_manager.ExistsIdentity(&newId)
	if err == nil {
		return userOtp, err
	}

	// copy default options and adjust name of new account
	otpUserOpts := GenerateOtpOpts
	otpUserOpts.AccountName = newId

	// generate otp key
	key, _ := totp.Generate(otpUserOpts)

	newUser := content_modifier.UserOtp{
		Id:  newId,
		Key: *key,
	}

	// save new identity
	err = id_manager.CreateIdentity(&newId, &newUser)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &newUser, nil
}
