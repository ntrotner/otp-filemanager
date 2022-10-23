package http_api

import (
	"html/template"
	"log"
	"net/http"
	otpresponder "otp-filemanager/http-api/otp-identity-responder"
	otp_login_responder "otp-filemanager/http-api/otp-login-responder"
	permissioncontroller "otp-filemanager/permission-controller"
	"time"
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
		// get username and password
		id, clientOverlappingCode, ok := r.BasicAuth()

		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		currentTime := time.Now()
		log.Println(currentTime)
		log.Println("Client:", id)

		// check if user exists and code is valid
		foundID, err := permissioncontroller.ChallengeLogin(&id, &clientOverlappingCode, &currentTime)

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("Access Denied"))
			log.Println("Login Failed", id)
			return
		}

		responder := otp_login_responder.SelectResponder(&mode, foundID, &w)
		responder.Send()
		log.Println("Login Successful", id)
	})
}
