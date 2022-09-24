package otp_login_responder

import "fmt"

type FileListResponder struct {
	Tool LoginResponderTool
}

type FileListTemplate struct {
	UserID string
}

func (r FileListResponder) Send() {
	(*r.Tool.HttpResponder).WriteHeader(200)
	(*r.Tool.HttpResponder).Header().Set("Content-Type", "application/json")
	(*r.Tool.HttpResponder).Write([]byte(fmt.Sprintln(r.Tool.Files)))
}
