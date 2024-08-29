package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"example.com/m/data"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (p *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
	case http.MethodPost:
		p.addProduct(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *ProductsHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(lp)
	if err != nil {
		http.Error(w, "oops something went wrong!", http.StatusBadRequest)
		return
	}

}

func (p *ProductsHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.l.Println("Error reading body:", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Received body: %s", body)

	prod := &data.Product{}
	err = prod.FromJSON(bytes.NewReader(body))
	if err != nil {
		p.l.Println("Unmarshal error:", err)
		http.Error(w, "Unable to unmarshal JSON "+err.Error(), http.StatusBadRequest)
		return
	}
	p.l.Printf("Product: %#v", prod)

	w.WriteHeader(http.StatusCreated)
}
