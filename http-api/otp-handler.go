package http_api

import (
	"bytes"
	"image/png"
	"net/http"
	permission_controller "otp-filemanager/permission-controller"
)

// initialize otp related endpoints
func OTPHandler() {
	// handle creation of a new identity
	http.HandleFunc("/createIdentity", func(w http.ResponseWriter, r *http.Request) {
		// create new random identity
		new_user, err := permission_controller.CreateIdentity()

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Couldn't create new Identity"))
			return
		}

		// create qr code image
		image, err := new_user.Key.Image(128*3, 128*3)

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

}
