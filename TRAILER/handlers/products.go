package handlers

import (
	"encoding/json"
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
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
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
