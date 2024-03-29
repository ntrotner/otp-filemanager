package id_manager

import (
	"errors"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
	file_system "otp-filemanager/entity-controller/id-manager/content-modifier/file-system"
)

var (
	ExistingIDs map[string]*content_modifier.UserOtp
	Modifier    content_modifier.Modifier
)

// InitializeIDManager prepares the data structures
func InitializeIDManager(modifierType *int8, expirationTime *int64) {
	switch *modifierType {
	case content_modifier.FileSystem:
		Modifier = file_system.CreateFileSystemModifier()
	default:
		Modifier = file_system.CreateFileSystemModifier()
	}

	Modifier.OtpModifier.InitializeOTPModifier()
	ExistingIDs = *Modifier.OtpModifier.ReadAllIdentities()
	OrchestrateExpirationCheck(expirationTime)
}

// ExistsIdentity check if identity exists
func ExistsIdentity(id *string) (*content_modifier.UserOtp, error) {
	user, existsUser := ExistingIDs[*id]

	if existsUser {
		return user, nil
	}

	return user, errors.New("user couldn't be found")
}

// CreateIdentity creates new identity in memory and filesystem
func CreateIdentity(id *string, user_otp *content_modifier.UserOtp) error {
	err := Modifier.OtpModifier.WriteIdentity(id, user_otp)

	if err != nil {
		return errors.New("couldn't write user to database")
	}

	ExistingIDs[*id] = user_otp

	return nil
}

// DeleteIdentity deletes an identity in memory and filesystem
func DeleteIdentity(id *string) error {
	err := Modifier.OtpModifier.DeleteIdentity(id)
	delete(ExistingIDs, *id)

	if err != nil {
		return errors.New("couldn't delete user from database")
	}

	return nil
}

func ReadFilesOfIdentity(id *string) []string {
	return Modifier.OtpModifier.ReadFilesOfIdentity(id)
}
