package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	contentIndex  = "./content/index.html"
	contentImages = "./content/images"

	tlsCrt = "./tls/server.crt"
	tlsKey = "./tls/server.key"

	port = ":8080"
)

var (
	tpl = template.Must(template.ParseFiles(contentIndex))
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			tpl.Execute(w, nil)
		},
	)

	r.HandleFunc(
		"/{path}",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", contentImages, mux.Vars(r)["path"]))
		},
	)

	log.Printf("service started on port %s", port)
	if err := http.ListenAndServeTLS(port, tlsCrt, tlsKey, r); err != nil {
		log.Printf("start server error: %s", err.Error())
	}
}
