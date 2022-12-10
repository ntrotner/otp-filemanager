package http_api

import (
	"html/template"
	"log"
	"net/http"
	"otp-filemanager/http-api/common"
	otpresponder "otp-filemanager/http-api/responder/otp-identity-responder"
	otp_login_responder "otp-filemanager/http-api/responder/otp-login-responder"
	permissioncontroller "otp-filemanager/permission-controller"
)

// initialize otp related endpoints
func OTPHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("http-api/html-templates/home.gohtml"))
		err := tmpl.Execute(w, nil)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Couldn't build HTML"))
			return
		}
	})

	// handle creation of a new identity
	http.HandleFunc("/createIdentity", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		// see otpresponder for all "resp" possibilities
		mode := otpresponder.IdentityResponse(query.Get("resp"))
		// create new random identity
		newUser, err := permissioncontroller.CreateIdentity()

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Couldn't create new Identity"))
			return
		}

		responder := otpresponder.SelectResponder(&mode, newUser, &w)
		responder.Send()
	})

	// handle login of existant user
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		// see loginresponder for all "resp" possibilities
		mode := otp_login_responder.LoginResponse(query.Get("resp"))

		// check if user exists and code is valid
		foundID, err := common.ChallengeLoginHTTP(r, w)

		if err != nil {
			return
		}

		responder := otp_login_responder.SelectResponder(&mode, foundID, &w)
		responder.Send()
		log.Println("Login Successful", foundID.Id)
	})
}
