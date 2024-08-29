package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/IamMaheshGurung/Microservices.git/TRAILER/handlers"
)

func main() {

	pl := log.New(os.Stdout, "courses-api", log.LstdFlags)
	hh := handlers.NewProductHandler(pl)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	server := http.Server{
		Addr:        ":9090",
		Handler:     sm,
		IdleTimeout: 130 * time.Second,
	}
	go func() {
		pl.Printf("Server starting at localhost:9090")
		err := server.ListenAndServe()
		if err != nil {
			pl.Printf("Error at serving in local host")
			os.Exit(1)
		}
	}()

	//For GraceFul Shutdown

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	pl.Printf("Requeestd for Gracful Termination %s", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)

}
