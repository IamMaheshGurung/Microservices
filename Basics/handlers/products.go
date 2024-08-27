package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/m/data"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}
func (p Product) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	encoder := json.NewEncoder(w)
	err := encoder.Encode(lp)
	if err != nil {
		http.Error(w, "oops something went wrong!", http.StatusBadRequest)
		return
	}
}
