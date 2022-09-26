package main

import (
	"fmt"
	"html/template"
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

	http.ListenAndServeTLS(port, tlsCrt, tlsKey, r)
}
