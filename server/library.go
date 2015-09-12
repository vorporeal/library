// library.go
package library

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	staticFileRoot = "./static/"
)

func init() {
	r := mux.NewRouter()
    serveSingle(r, "/", "app/index.html")

	http.Handle("/", r)
}

func redirect(r *mux.Router, route, destination string) {
  	r.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
    	http.Redirect(w, r, destination, http.StatusFound)
  	})
}

func serveSingle(r *mux.Router, route, filename string) {
	r.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
