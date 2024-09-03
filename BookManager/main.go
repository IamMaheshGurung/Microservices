package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"xample.com/m/handlers"
)

func main() {

	l := log.New(os.Stdout, "Books-api", log.LstdFlags)

	bh := handlers.NewBook(l)
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/book", bh.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", bh.GetBookId).Methods(http.MethodGet)
	router.HandleFunc("/book", bh.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", bh.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", bh.DeleteBook).Methods(http.MethodDelete)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	server := &http.Server{
		Addr:        ":8080",
		Handler:     c.Handler(router),
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		fmt.Printf("Server Loading at %#v", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {

			log.Fatalf("Couldnot listen on %s: %v\n ", server.Addr, err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sg := <-sigChan
	log.Printf("Rrecived signal %v to shutdown server ", sg)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(tc); err != nil {
		log.Fatalf("Server Forced to shutdown: %v", err)

	}
	log.Println("Server exiting")

}
