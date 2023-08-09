package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Server is running on port 7070")
	if err := http.ListenAndServe(":7070", nil); err != nil {
		log.Fatal(err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website!")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Age= %v\n", age)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello World!")
}
