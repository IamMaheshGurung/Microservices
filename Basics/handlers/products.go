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

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	w.Header().Set("Conetnt-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(lp)
	if err != nil {
		http.Error(w, "oops something went wrong!", http.StatusBadRequest)
		return
	}
}

func (p *Product) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)

	}
	p.l.Printf("Prod: %#v", prod)
}
