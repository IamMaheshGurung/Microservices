package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	server := http.Server{
		Addr:        ":8080",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		fmt.Println("Server running locally")
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Server error %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, Cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer Cancel()

	fmt.Println("Shutting down the server....")
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown error %v\n", err)
	}
	fmt.Println("Server shut down gracefully.....")

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("index.html", "nav.html", "home.html", "form.html")
		if err != nil {
			http.Error(w, "Unable to load the template", http.StatusInternalServerError)
			fmt.Printf("%v\n", err)

			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
			fmt.Printf("%v\n", err)

		}
	}

}
