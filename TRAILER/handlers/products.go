package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/IamMaheshGurung/Microservices.git/TRAILER/data"
)

type ProductHandler struct {
	pl *log.Logger
}

func NewProductHandler(pl *log.Logger) *ProductHandler {
	return &ProductHandler{pl}
}

func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		p.getProducts(w, r)
	case http.MethodPost:
		p.addProducts(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	jp := data.GetProducts()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(jp)
	if err != nil {
		http.Error(w, "Opps Spmething went very wrong", http.StatusBadRequest)
		return

	}
}

func (p *ProductHandler) addProducts(w http.ResponseWriter, r *http.Request) {
	p.pl.Println("Handle post product")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.pl.Println("Error in reading Body", err)
		http.Error(w, "Unable to ready the body", http.StatusInternalServerError)
		return
	}
	prod := &data.Product{}
	err = prod.FromJSON(bytes.NewReader(body))
	if err != nil {
		p.pl.Println("Unmarshal error", err)
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	p.pl.Printf("Product been added %#v", prod)

}
