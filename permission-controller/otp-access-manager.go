package permission_controller

import (
	"errors"
	content_modifier "otp-filemanager/content-modifier"
	idmanager "otp-filemanager/permission-controller/id-manager"
	"time"

	"github.com/pquerna/otp/totp"
)

// ChallengeLogin finds user and checks if code is valid
func ChallengeLogin(id *string, clientCode *string, time *time.Time) (*content_modifier.UserOtp, error) {
	foundID, err := idmanager.ExistsIdentity(id)

	if err != nil {
		return foundID, err
	}

	validPassword, _ := totp.ValidateCustom(*clientCode, foundID.Key.Secret(), *time, ValidateOtpOpts)

	if !validPassword {
		return foundID, errors.New("invalid Password")
	}

	return foundID, nil
}

func ChallengeReadFile() {
}

func ChallengeWriteFile() {
}

func ChallengeDeleteFile() {
}

func ChallengeDeleteIdentity() {
}
