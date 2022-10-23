package otp_login_responder

import (
	"net/http"
	content_modifier "otp-filemanager/content-modifier"
)

func SelectResponder(mode *LoginResponse, foundID *content_modifier.UserOtp, w *http.ResponseWriter) LoginResponder {
	files := content_modifier.ReadFilesOfIdentity(&foundID.Id)
	responderTool := LoginResponderTool{User: foundID, HttpResponder: w, Files: files}

	switch *mode {
	case AccountInterface:
		return AccountInterfaceResponder{Tool: responderTool}
	default:
		return FileListResponder{Tool: responderTool}
	}
}
