package otp_identity_responder

type UrlResponder struct {
	Tool IdentityResponderTool
}

func (r UrlResponder) Send() {
	// create url string
	url := r.Tool.NewUser.Key.URL()

	(*r.Tool.HttpResponder).Header().Set("Content-Type", "text/plain")
	(*r.Tool.HttpResponder).WriteHeader(200)
	(*r.Tool.HttpResponder).Write([]byte(url))
}
