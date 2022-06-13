package otp_identity_responder

import (
	"bytes"
	"image/png"
)

type QRCodeResponder struct {
	Tool IdentityResponderTool
}

func (r QRCodeResponder) Send() {
	// create qr code image
	image, err := r.Tool.NewUser.Key.Image(128*3, 128*3)

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("Couldn't show QR code"))
		return
	}

	// convert image to bytes in a buffer
	buf := new(bytes.Buffer)
	err = png.Encode(buf, image)

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("QR code encoding failed"))
		return
	}

	// send bytes
	(*r.Tool.HttpResponder).Header().Set("Content-Type", "image/png")
	(*r.Tool.HttpResponder).Write(buf.Bytes())
}
