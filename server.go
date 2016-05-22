package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const version = "0.0.1"

func main() {
	fmt.Printf("Starting DisclosureJockey server...\n")
	fmt.Printf("Version v%s\n", version)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", <-guids)
}
