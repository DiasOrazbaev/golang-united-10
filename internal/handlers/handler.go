package handlers

import "net/http"

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}
