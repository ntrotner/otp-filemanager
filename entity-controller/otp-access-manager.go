package permission_controller

import (
	"errors"
	"log"
	"mime/multipart"
	id_manager "otp-filemanager/entity-controller/id-manager"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
	"time"

	"github.com/pquerna/otp/totp"
)

// ChallengeLogin finds user and checks if code is valid
func ChallengeLogin(id *string, clientCode *string, time *time.Time) (*content_modifier.UserOtp, error) {
	foundID, err := id_manager.ExistsIdentity(id)

	if err != nil {
		return foundID, err
	}

	validPassword, _ := totp.ValidateCustom(*clientCode, foundID.Key.Secret(), *time, ValidateOtpOpts)

	if !validPassword {
		return foundID, errors.New("invalid Password")
	}

	return foundID, nil
}

func ChallengeReadFile(id *string, clientCode *string, time *time.Time, fileName *string) (*[]byte, error) {
	foundID, err := ChallengeLogin(id, clientCode, time)
	if err != nil {
		log.Println("Error for Read File:", err)
		return nil, err
	}

	file, err := id_manager.Modifier.FileModifier.ReadFile(&foundID.Id, fileName)
	if err != nil {
		log.Println("Error for Read File:", err)
		return nil, err
	}

	return file, nil
}

func ChallengeWriteFile(id *string, clientCode *string, time *time.Time, fileName *string, file *multipart.File) error {
	foundID, err := ChallengeLogin(id, clientCode, time)
	if err != nil {
		return err
	}
	err = id_manager.Modifier.FileModifier.WriteFile(&foundID.Id, fileName, file)

	return err
}

func ChallengeDeleteFile(id *string, clientCode *string, time *time.Time, fileName *string) error {
	foundID, err := ChallengeLogin(id, clientCode, time)
	if err != nil {
		log.Println("Error for Delete File:", err)
		return err
	}

	err = id_manager.Modifier.FileModifier.DeleteFile(&foundID.Id, fileName)
	if err != nil {
		log.Println("Error for Delete Identity:", err)
		return err
	}

	return nil
}

func ChallengeDeleteIdentity(id *string, clientCode *string, time *time.Time) error {
	foundID, err := ChallengeLogin(id, clientCode, time)
	if err != nil {
		log.Println("Error for Delete Identity:", err)
		return err
	}

	err = id_manager.DeleteIdentity(&foundID.Id)
	if err != nil {
		log.Println("Error for Delete Identity:", err)
		return err
	}

	return nil
}
