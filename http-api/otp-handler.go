package http_api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	content_modifier "otp-filemanager/content-modifier"
	"otp-filemanager/helper"
	otpresponder "otp-filemanager/http-api/otp-identity-responder"
	permissioncontroller "otp-filemanager/permission-controller"
	"path"
	"time"
)

// initialize otp related endpoints
func OTPHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html-templates/home.gohtml"))
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

		var responder otpresponder.IdentityResponder
		responderTool := otpresponder.IdentityResponderTool{NewUser: newUser, HttpResponder: &w}

		// select mode of response by the use of "resp", whereas default is plain text
		switch mode {
		case otpresponder.QR:
			responder = otpresponder.QRCodeResponder{Tool: responderTool}
			break
		case otpresponder.Combination:
			responder = otpresponder.CombinedResponder{Tool: responderTool}
			break
		case otpresponder.Url:
		default:
			responder = otpresponder.UrlResponder{Tool: responderTool}
		}

		responder.Send()
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// get username and password
		id, clientOverlappingCode, ok := r.BasicAuth()

		if ok {
			currentTime := time.Now()
			log.Println("Client:", id)

			// check if user exists and code is valid
			foundID, err := permissioncontroller.ChallengeLogin(&id, &clientOverlappingCode, &currentTime)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte("Access Denied"))
				log.Println("Login Failed", id)
				return
			} else {
				pathToFilesOfFoundID := path.Join(content_modifier.PathToFilesOfIdentities, foundID.Id)
				files := helper.ReadFileNamesOfDirectory(&pathToFilesOfFoundID)

				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(fmt.Sprintln(files)))
				log.Println("Login Successful", id)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
