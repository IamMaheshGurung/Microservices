package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", Homepage)
	http.ListenAndServe(":8080", nil)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "htmx app",
		Message: "I am from golang",
	}
	tmpl.Execute(w, data)
}
