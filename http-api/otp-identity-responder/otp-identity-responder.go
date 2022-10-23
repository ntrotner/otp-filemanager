package otp_identity_responder

import (
	"net/http"
	content_modifier "otp-filemanager/content-modifier"
)

func SelectResponder(mode *IdentityResponse, newUser *content_modifier.UserOtp, w *http.ResponseWriter) IdentityResponder {
	responderTool := IdentityResponderTool{NewUser: newUser, HttpResponder: w}

	// select mode of response by the use of "resp", whereas default is plain text
	switch *mode {
	case QR:
		return QRCodeResponder{Tool: responderTool}
	case Combination:
		return CombinedResponder{Tool: responderTool}
	default:
		return UrlResponder{Tool: responderTool}
	}
}
