package http_api

import (
	"log"
	"net/http"
)

func InitializeHTTPServer(port *string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	log.Println("Listening on Port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
