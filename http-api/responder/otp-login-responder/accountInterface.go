package otp_login_responder

import (
	"html/template"
)

type AccountInterfaceResponder struct {
	Tool LoginResponderTool
}

type AccountInterfaceTemplate struct {
	UserID   string
	Files    []string
	Password string
}

func (r AccountInterfaceResponder) Send() {
	tmpl := template.Must(template.ParseFiles("http-api/html-templates/accountInterface.gohtml"))
	data := AccountInterfaceTemplate{UserID: r.Tool.User.Id, Files: r.Tool.Files, Password: (*r.Tool.HttpResponder).Header().Get("Authorization")}
	err := tmpl.Execute(*r.Tool.HttpResponder, data)

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("Couldn't build HTML"))
		return
	}
}
