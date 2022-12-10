package otp_identity_responder

import (
	"net/http"
	content_modifier "otp-filemanager/permission-controller/id-manager/content-modifier"
)

type IdentityResponse string

const (
	Combination IdentityResponse = "cb"
	QR          IdentityResponse = "qr"
	Url         IdentityResponse = "url"
)

type IdentityResponder interface {
	// Send handles the event of sending the response to the client
	// while also handling error events
	Send()
}

type IdentityResponderTool struct {
	NewUser       *content_modifier.UserOtp
	HttpResponder *http.ResponseWriter
}
