package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hi I am from the server")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server loading at local host")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Unable to load the server")
	}

}
