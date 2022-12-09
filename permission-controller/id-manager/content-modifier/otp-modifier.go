package content_modifier

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"otp-filemanager/helper"
	"path"
	"path/filepath"
	"time"

	"github.com/pquerna/otp"
)

// InitializeOTPModifier prepares the directory for saving new identities
func InitializeOTPModifier() {
	err := helper.CreateDirectory(PathToIdentities)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = helper.CreateDirectory(PathToFilesOfIdentities)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// ReadAllIdentities parses all identities in the directory
func ReadAllIdentities() *map[string]*UserOtp {
	var otpIdentities = make(map[string]*UserOtp)
	fileInfos, err := ioutil.ReadDir(PathToIdentities)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, element := range fileInfos {
		userId := helper.FilenameWithoutExtension(element.Name())
		parsedUser, err := ReadIdentity(&userId)

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
	byteFile, err := ioutil.ReadFile(filepath.Join(PathToIdentities, *id+".json"))

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

	parsedUser.Key = *otpKey
	parsedUser.Id = *id

	return &parsedUser, nil
}

// WriteIdentity creates a new identity file
func WriteIdentity(id *string, identity *UserOtp) error {
	fsUser := FilesystemUserOtp{
		URL_Key:     identity.Key.URL(),
		Issued_Date: time.Now().String(),
	}

	file, err := json.MarshalIndent(fsUser, "", " ")

	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(filepath.Join(PathToIdentities, *id+".json"), file, 0644)

	if err != nil {
		log.Println(err)
		return err
	}

	err = os.MkdirAll(filepath.Join(PathToFilesOfIdentities, *id), 0755)

	if err != nil {
		log.Println(err)
		os.Remove(filepath.Join(PathToIdentities, *id+".json"))
		return err
	}

	_, err = ReadIdentity(id)
	return err
}

// DeleteIdentity deletes the file for the identity
func DeleteIdentity(id *string) error {
	err := os.Remove(filepath.Join(PathToIdentities, *id+".json"))
	if err != nil {
		log.Println(err)
	}

	err = os.Remove(filepath.Join(PathToFilesOfIdentities, *id))
	if err != nil {
		log.Println(err)
	}

	return err
}

// ReadFilesOfIdentity gets list of files related to an id
func ReadFilesOfIdentity(id *string) []string {
	pathToFilesOfFoundID := path.Join(PathToFilesOfIdentities, *id)
	return helper.ReadFileNamesOfDirectory(&pathToFilesOfFoundID)
}
