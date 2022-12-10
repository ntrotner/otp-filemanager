package otp_identity_responder

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image/png"
)

type CombinedResponder struct {
	Tool IdentityResponderTool
}

type CombinedTemplate struct {
	OtpUrl template.URL
	Image  string
}

func (r CombinedResponder) Send() {
	// create url string
	url := r.Tool.NewUser.Key.URL()

	// create qr code image
	image, err := r.Tool.NewUser.Key.Image(128*3, 128*3)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, image)

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("Couldn't show QR code"))
		return
	}

	tmpl := template.Must(template.ParseFiles("http-api/html-templates/combinedIdentity.gohtml"))
	err = tmpl.Execute(*r.Tool.HttpResponder, CombinedTemplate{OtpUrl: template.URL(url), Image: base64.RawStdEncoding.EncodeToString(buf.Bytes())})

	if err != nil {
		(*r.Tool.HttpResponder).WriteHeader(500)
		(*r.Tool.HttpResponder).Write([]byte("Couldn't build HTML"))
		return
	}
}
