package http_api

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	permissioncontroller "otp-filemanager/permission-controller"
	idmanager "otp-filemanager/permission-controller/id-manager"
	"time"

	"github.com/pquerna/otp/totp"
)

// initialize otp related endpoints
func OTPHandler() {
	// handle creation of a new identity
	http.HandleFunc("/createIdentity", func(w http.ResponseWriter, r *http.Request) {
		// create new random identity
		newUser, err := permissioncontroller.CreateIdentity()

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Couldn't create new Identity"))
			return
		}

		// create qr code image
		image, err := newUser.Key.Image(128*3, 128*3)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Couldn't show QR code"))
			return
		}

		// convert image to bytes in a buffer
		buf := new(bytes.Buffer)
		err = png.Encode(buf, image)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("QR code encoding failed"))
			return
		}

		// send bytes
		w.Header().Set("Content-Type", "image/png")
		w.Write(buf.Bytes())
	})

	http.HandleFunc("/challengeLogin", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		// get username and password
		id := query.Get("id")
		clientOverlappingCode := query.Get("otp")

		log.Println("Client:", id)

		foundID, err := idmanager.ExistsIdentity(&id)

		if err != nil {
			w.WriteHeader(500)
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
			w.WriteHeader(400)
			w.Write([]byte("Access Denied"))
			log.Println("Login Failed", id)
		}
	})
}
