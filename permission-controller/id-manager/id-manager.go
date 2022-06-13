package id_manager

import (
	"errors"
	content_modifier "otp-filemanager/content-modifier"
)

var (
	existingIDs map[string]*content_modifier.UserOtp
)

// InitializeIDManager prepares the data structures
func InitializeIDManager() {
	existingIDs = *content_modifier.ReadAllIdentities()
}

// ExistsIdentity check if identity exists
func ExistsIdentity(id *string) (*content_modifier.UserOtp, error) {
	user, existsUser := existingIDs[*id]

	if existsUser {
		return user, nil
	}

	return user, errors.New("User couldn't be found")
}

// CreateIdentity creates new identity in memory and filesystem
func CreateIdentity(id *string, user_otp *content_modifier.UserOtp) error {
	err := content_modifier.WriteIdentity(id, user_otp)

	if err != nil {
		return errors.New("Couldn't write user to database")
	}

	existingIDs[*id] = user_otp

	return nil
}
