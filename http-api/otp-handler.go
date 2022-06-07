package http_api

import (
	"fmt"
	"log"
	"net/http"
	otpresponder "otp-filemanager/http-api/otp-identity-responder"
	permissioncontroller "otp-filemanager/permission-controller"
	idmanager "otp-filemanager/permission-controller/id-manager"
	"time"

	"github.com/pquerna/otp/totp"
)

// initialize otp related endpoints
func OTPHandler() {
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
			log.Println("Client:", id)

			foundID, err := idmanager.ExistsIdentity(&id)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte("Identity wasn't found"))
				log.Println("Invalid Identity", id)
				return
			}
			validCode, err := totp.ValidateCustom(clientOverlappingCode, foundID.Key.Secret(), time.Now(), permissioncontroller.ValidateOtpOpts)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("Identity unresponsive"))
				log.Println("Identity unresponsive", id)
				log.Println(err)
				return
			}

			if validCode {
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(fmt.Sprintln(foundID.Files)))
				log.Println("Login Successful", id)
			} else {
				w.WriteHeader(401)
				w.Write([]byte("Access Denied"))
				log.Println("Login Failed", id)
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
