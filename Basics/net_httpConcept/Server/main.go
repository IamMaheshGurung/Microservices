package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := http.NewServeMux()
	home := http.HandlerFunc(handler)
	form := http.HandlerFunc(formHandler)
	router.Handle("/", home)
	router.Handle("/form", form)

	server := http.Server{
		Addr:        ":8080",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		fmt.Println("Local host running at :8080")
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server Error:%v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	println("Gracefully shutting down the server")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("SHutdown due to %v", err)

	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, err := fmt.Println("Hi Everyone I am the homepage")
		if err != nil {
			http.Error(w, "Unable to get the webpage", http.StatusBadRequest)
		}

	}
	http.Error(w, "Unable to get the webpage", http.StatusBadRequest)

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "unable to parse form", http.StatusBadRequest)
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		response := fmt.Sprintf("Data Received %s as name and %s as email", name, email)
		fmt.Fprintln(w, response)
	} else {
		http.Error(w, "Incorrect request method", http.StatusBadRequest)
	}

}
