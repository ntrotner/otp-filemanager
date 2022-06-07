package otp_identity_responder

import (
	"net/http"
	content_modifier "otp-filemanager/content-modifier"
)

type IdentityResponse string

const (
	Combination IdentityResponse = "cb"
	QR          IdentityResponse = "qr"
	Url         IdentityResponse = "url"
)

type IdentityResponder interface {
	Send()
}

type IdentityResponderTool struct {
	NewUser       *content_modifier.UserOtp
	HttpResponder *http.ResponseWriter
}
