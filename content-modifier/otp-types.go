package content_modifier

import (
	"github.com/pquerna/otp"
)

type UserOtp struct {
	Id    string
	Key   otp.Key
	Files []string
}

type FilesystemUserOtp struct {
	URL_Key     string   `json:"url"`
	Issued_Date string   `json:"date"`
	Files       []string `json:"files"`
}

type LoginChallenge struct {
	Files []string `json:"files"`
}
