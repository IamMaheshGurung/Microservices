package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/IamMaheshGurung/htmxtailwind/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "Tweets-API", log.LstdFlags)
	th := handlers.NewTweet(l)

	ls := mux.NewRouter()
	ls.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../Front"))))
	ls.HandleFunc("/tweet", th.AddTweets).Methods(http.MethodPost)
	ls.HandleFunc("/tweet-list", th.GetTweet).Methods(http.MethodGet)

	server := http.Server{
		Addr:        ":8080",
		Handler:     ls,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		println("Running at local host")
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
		log.Fatalf("Server forced to shut down %s", err)
	}
	log.Println("Gracefully stopped")
}
