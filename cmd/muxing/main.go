package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Start(host string, port int) {
	router := mux.NewRouter()

	log.Println(fmt.Sprintf("Starting API server on %s:%d", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	host := os.Getenv("Host")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
