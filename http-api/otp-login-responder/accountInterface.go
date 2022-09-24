package otp_login_responder

import (
	"html/template"
)

type AccountInterfaceResponder struct {
	Tool LoginResponderTool
}

type AccountInterfaceTemplate struct {
	UserID string
	Files  []string
}

func (r AccountInterfaceResponder) Send() {
	tmpl := template.Must(template.ParseFiles("html-templates/accountInterface.gohtml"))
	err := tmpl.Execute(*r.Tool.HttpResponder, AccountInterfaceTemplate{UserID: r.Tool.User.Id, Files: r.Tool.Files})

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("Couldn't build HTML"))
		return
	}
}
