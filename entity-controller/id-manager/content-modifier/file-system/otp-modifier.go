package file_system

import (
	"encoding/json"
	"log"
	"os"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
	"otp-filemanager/entity-controller/security"
	"otp-filemanager/helper"
	"path"
	"path/filepath"
	"time"

	"github.com/pquerna/otp"
)

const DateFormat = time.UnixDate

var PathToIdentities = filepath.Join("live-data", "identities")

var PathToFilesOfIdentities = filepath.Join("live-data", "files")

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
func ReadAllIdentities() *map[string]*content_modifier.UserOtp {
	var otpIdentities = make(map[string]*content_modifier.UserOtp)
	fileInfos, err := os.ReadDir(PathToIdentities)
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
func ReadIdentity(id *string) (*content_modifier.UserOtp, error) {
	readUser := content_modifier.FilesystemUserOtp{}
	parsedUser := content_modifier.UserOtp{}
	byteFile, err := os.ReadFile(filepath.Join(PathToIdentities, *id+".json"))

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	decryptedIdentity, err := security.Decrypt(&byteFile)
	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	err = json.Unmarshal(*decryptedIdentity, &readUser)

	if err != nil {
		log.Println("Couldn't parse JSON for Identity")
		return &parsedUser, err
	}

	otpKey, err := otp.NewKeyFromURL(readUser.URL_Key)

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	parsedDate, err := time.Parse(DateFormat, readUser.Issued_Date)

	if err != nil {
		log.Println(err)
		return &parsedUser, err
	}

	parsedUser.IssuedDate = parsedDate
	parsedUser.Key = *otpKey
	parsedUser.Id = *id

	return &parsedUser, nil
}

// WriteIdentity creates a new identity file
func WriteIdentity(id *string, identity *content_modifier.UserOtp) error {
	fsUser := content_modifier.FilesystemUserOtp{
		URL_Key:     identity.Key.URL(),
		Issued_Date: time.Now().Format(DateFormat),
	}

	file, err := json.MarshalIndent(fsUser, "", " ")

	if err != nil {
		log.Println(err)
		return err
	}

	encryptedIdentity, err := security.Encrypt(&file)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(filepath.Join(PathToIdentities, *id+".json"), *encryptedIdentity, 0644)

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

	readIdentity, err := ReadIdentity(id)
	*identity = *readIdentity

	return err
}

// DeleteIdentity deletes the file for the identity
func DeleteIdentity(id *string) error {
	err := os.RemoveAll(filepath.Join(PathToIdentities, *id+".json"))
	if err != nil {
		log.Println(err)
	}

	err = os.RemoveAll(filepath.Join(PathToFilesOfIdentities, *id))
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
