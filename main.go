package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(w, "POST request is successful")
	name := r.FormValue("name")
	company := r.FormValue("company")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "COmapny = %s\n", company)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "OTHER METHODS NOT SUPPORTED", http.StatusNotFound)
	}

	fmt.Fprintf(w, "WELCOME TO PATH: /hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("STARTING SERVER AAT PORT 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
