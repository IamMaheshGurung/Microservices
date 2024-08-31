package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xample.com/m/handlers"
)

func main() {
	lg := log.New(os.Stdout, "Course-Api", log.LstdFlags)

	hg := handlers.NewCourse(lg)

	mux := http.NewServeMux()

	mux.Handle("/", hg)

	server := http.Server{
		Addr:        ":9090",
		Handler:     mux,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		lg.Println("Server startng now at local host 9090:")
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Unable to connect Server")
		}
	}()

	//GraceFul shutdown

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)
	sg := <-sigChan
	lg.Println("Requested for the graceful shutdown: ", sg)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tc)

}
