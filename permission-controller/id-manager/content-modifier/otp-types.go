package content_modifier

import (
	"github.com/pquerna/otp"
)

const FileSystem int8 = 0

type Modifier struct {
	FileModifier FileModifier
	OtpModifier  OtpModifier
}

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
