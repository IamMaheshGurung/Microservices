package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/m/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//handler
	hh := handlers.NewProducts(l)

	//http.Handle("/", hh)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	//sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:        ":8080",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}() // Add closing parenthesis here

	//graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sg := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sg)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
