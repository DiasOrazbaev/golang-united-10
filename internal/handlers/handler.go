package handlers

import (
	"fmt"
	"io"
	"log"
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

func DataBodyParamHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Fail: Parse body")
	}
	w.Write([]byte("I got message:\n" + string(body)))
}
