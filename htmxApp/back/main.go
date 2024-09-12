package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/back/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../front"))))
	r.HandleFunc("/message", handlers.GetMessage).Methods(http.MethodGet)
	r.HandleFunc("/submit", handlers.Submithandler).Methods(http.MethodPost)

	server := http.Server{
		Addr:        ":8080",
		Handler:     r,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		log.Println("Running at local host")
		if err := server.ListenAndServe(); err != nil {
			println("Server Error,", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown %s", err)
	}
	log.Println("Gracefully shutdown")
}
