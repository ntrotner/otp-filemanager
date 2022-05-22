package permission_controller

import (
	content_modifier "otp-filemanager/content-modifier"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/teris-io/shortid"
)

var (
	id_generator *shortid.Shortid
	otp_opts     totp.GenerateOpts
)

// create worker for generating identities and default options for otp
func InitializeOTPGenerator(seed *uint64, issuer *string, period *uint) {
	id_generator, _ = shortid.New(1, shortid.DefaultABC, *seed)

	otp_opts = totp.GenerateOpts{
		Issuer:      *issuer,
		AccountName: "",
		Period:      *period,
		SecretSize:  0,
		Secret:      nil,
		Digits:      otp.DigitsEight,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        nil,
	}
}

// create new random identity
func CreateIdentity() (*content_modifier.User_otp, error) {
	// create new identity
	new_id, err := id_generator.Generate()

	if err != nil {
		return nil, err
	}

	// copy default options and adjust name of new account
	otp_user_opts := otp_opts
	otp_user_opts.AccountName = new_id

	// generate otp key
	key, _ := totp.Generate(otp_user_opts)

	new_user := content_modifier.User_otp{
		Id:  new_id,
		Key: *key,
	}

	// save new identity
	err = content_modifier.WriteIdentity(&new_user)

	if err != nil {
		return nil, err
	}

	return &new_user, nil
}
