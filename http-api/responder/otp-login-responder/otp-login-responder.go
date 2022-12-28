package otp_login_responder

import (
	"net/http"
	id_manager "otp-filemanager/entity-controller/id-manager"
	content_modifier "otp-filemanager/entity-controller/id-manager/content-modifier"
)

func SelectResponder(mode *LoginResponse, foundID *content_modifier.UserOtp, w *http.ResponseWriter) LoginResponder {
	files := id_manager.ReadFilesOfIdentity(&foundID.Id)
	responderTool := LoginResponderTool{User: foundID, HttpResponder: w, Files: files}

	switch *mode {
	case AccountInterface:
		return AccountInterfaceResponder{Tool: responderTool}
	default:
		return FileListResponder{Tool: responderTool}
	}
}
