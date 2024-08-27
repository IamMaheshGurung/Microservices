package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"example.com/m/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//handler
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	//http.Handle("/", hh)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:        ":8080",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
	}
	s.ListenAndServe()

	//graceful shutdown

}
