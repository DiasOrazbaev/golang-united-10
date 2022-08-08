package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func NameParametrHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte(fmt.Sprintf("Hello, %v!", name)))
}
