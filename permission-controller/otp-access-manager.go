package permission_controller

import (
	"errors"
	"github.com/pquerna/otp/totp"
	content_modifier "otp-filemanager/content-modifier"
	idmanager "otp-filemanager/permission-controller/id-manager"
	"time"
)

// ChallengeLogin finds user and checks if code is valid
func ChallengeLogin(id *string, clientCode *string, time *time.Time) (*content_modifier.UserOtp, error) {
	foundID, err := idmanager.ExistsIdentity(id)

	if err != nil {
		return foundID, err
	}

	validPassword, err := totp.ValidateCustom(*clientCode, foundID.Key.Secret(), *time, ValidateOtpOpts)

	if validPassword {
		return foundID, errors.New("invalid Password")
	}

	return foundID, err
}

func ChallengeReadFile() {
}

func ChallengeWriteFile() {
}

func ChallengeDeleteFile() {
}

func ChallengeDeleteIdentity() {
}
