package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

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

func SumOfHeadersHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Println("Could not parse from header a")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Println("Could not parse from header a")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(a, b)

	w.Header().Add("a+b", strconv.FormatInt(int64(a+b), 10))
}
