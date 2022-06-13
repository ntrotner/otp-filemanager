package permission_controller

import (
	"log"
	contentmodifier "otp-filemanager/content-modifier"
	idmanager "otp-filemanager/permission-controller/id-manager"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/teris-io/shortid"
)

var (
	idGenerator     *shortid.Shortid
	GenerateOtpOpts totp.GenerateOpts
	ValidateOtpOpts totp.ValidateOpts
)

// InitializeOTPGenerator create worker for generating identities and default options for otp
func InitializeOTPGenerator(seed *uint64, issuer *string, period *uint) {
	idmanager.InitializeIDManager()
	idGenerator, _ = shortid.New(1, shortid.DefaultABC, *seed)

	GenerateOtpOpts = totp.GenerateOpts{
		Issuer:      *issuer,
		AccountName: "",
		Period:      *period,
		SecretSize:  0,
		Secret:      nil,
		Digits:      otp.DigitsEight,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        nil,
	}

	ValidateOtpOpts = totp.ValidateOpts{
		Period:    *period,
		Digits:    otp.DigitsEight,
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
