package otp_login_responder

import (
	"net/http"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
)

type LoginResponse string

const (
	FileList         LoginResponse = "fl"
	AccountInterface LoginResponse = "ai"
)

type LoginResponder interface {
	// Send handles the event of sending the response to the client
	// while also handling error events
	Send()
}

type LoginResponderTool struct {
	User          *content_modifier.UserOtp
	Files         []string
	HttpResponder *http.ResponseWriter
}
