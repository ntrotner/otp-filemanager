package content_modifier

import (
	"path/filepath"

	"github.com/pquerna/otp"
)

var PathToIdentities = filepath.Join("live-data", "identities")

var PathToFilesOfIdentities = filepath.Join("live-data", "files")

type UserOtp struct {
	Id  string
	Key otp.Key
}

type FilesystemUserOtp struct {
	URL_Key     string `json:"url"`
	Issued_Date string `json:"date"`
}

type LoginChallenge struct {
	Files []string `json:"files"`
}
