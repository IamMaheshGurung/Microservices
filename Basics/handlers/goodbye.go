package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("GoodBye!"))
	if err != nil {
		http.Error(w, "oops something went wrong!", http.StatusBadRequest)
		return
	}
}
