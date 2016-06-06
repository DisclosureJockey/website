package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const version = "0.0.1"

func main() {
	fmt.Printf("Starting DisclosureJockey server...\n")
	fmt.Printf("Version v%s\n", version)

	http.HandleFunc("/pages/", includeHandler)
	http.HandleFunc("/css/", includeHandler)
	//http.HandleFunc("/js", includeHandler)
	http.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Only use for files you want directly served.
func includeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[+] Include handler called.")
	filename := r.URL.Path[1:]
	log.Println("Probably GET", filename, r)
	http.ServeFile(w, r, filename)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", <-guids)
}
