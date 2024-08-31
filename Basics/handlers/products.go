package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	case http.MethodPut:
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println("Captured groups: ", g)
		if len(g) != 1 {
			p.l.Println("Invalid Id more than one url")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			p.l.Println("More than one capture")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URL unable to convert to number, id string ")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.l.Println("Got Id", idInt)

		p.updateProduct(idInt, w, r)

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
	data.AddProduct(prod)

	w.WriteHeader(http.StatusCreated)
}
func (p *ProductsHandler) updateProduct(idInt int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Marshal the Json ", http.StatusInternalServerError)
		return
	}
	err = data.UpdateProduct(idInt, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
