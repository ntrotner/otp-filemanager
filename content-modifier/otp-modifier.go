package content_modifier

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"otp-filemanager/helper"
	"path/filepath"
	"time"

	"github.com/pquerna/otp"
)

var pathToIdentities = filepath.Join("live-data", "identities")

// InitializeOTPModifier prepares the directory for saving new identities
func InitializeOTPModifier() {
	err := helper.CreateDirectory(pathToIdentities)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// ReadAllIdentities parses all identities in the directory
func ReadAllIdentities() *map[string]*UserOtp {
	var otpIdentities = make(map[string]*UserOtp)
	fileInfos, err := ioutil.ReadDir(pathToIdentities)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, element := range fileInfos {
		userId := helper.FilenameWithoutExtension(element.Name())
		parsedUser, err := ReadIdentity(&userId)
		log.Println(parsedUser)

		if err != nil {
			continue
		}

		otpIdentities[userId] = parsedUser
	}

	return &otpIdentities
}

// ReadIdentity reads a single identity
func ReadIdentity(id *string) (*UserOtp, error) {
	readUser := FilesystemUserOtp{}
	parsedUser := UserOtp{}
	byteFile, err := ioutil.ReadFile(filepath.Join(pathToIdentities, *id+".json"))

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	err = json.Unmarshal(byteFile, &readUser)

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	otpKey, err := otp.NewKeyFromURL(readUser.URL_Key)

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	log.Println(otpKey.URL())

	parsedUser.Files = readUser.Files
	parsedUser.Key = *otpKey
	parsedUser.Id = *id

	return &parsedUser, nil
}

// WriteIdentity creates a new identity file
func WriteIdentity(id *string, identity *UserOtp) error {
	fsUser := FilesystemUserOtp{
		URL_Key:     identity.Key.URL(),
		Issued_Date: time.Now().String(),
		Files:       identity.Files,
	}

	file, err := json.MarshalIndent(fsUser, "", " ")

	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(filepath.Join(pathToIdentities, *id+".json"), file, 0644)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ReadIdentity(id)
	return err
}

// DeleteIdentity deletes the file for the identity
func DeleteIdentity(id *string) error {
	err := os.Remove(filepath.Join(pathToIdentities, *id+".json"))

	if err != nil {
		log.Println(err)
	}
	return err
}
