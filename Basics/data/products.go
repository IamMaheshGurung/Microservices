package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	DeletedOn   time.Time `json:"deleted_on"`
}

// FromJSON decodes a JSON payload into the Product struct.
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// GetProducts returns a list of products.
func GetProducts() []*Product {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func getNextID() int {
	lp := productList[len(productList)-1]
	lp.ID++
	return lp.ID
}

// A sample product list with unique IDs
var productList = []*Product{
	{
		ID:          1,
		Name:        "Chayote Squash",
		Description: "A fresh and healthy chayote squash",
		Price:       1.99,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
		DeletedOn:   time.Time{}, // zero value for time.Time
	},
	{
		ID:          2, // Changed to unique ID
		Name:        "Ghiraula",
		Description: "A fresh and healthy Himalayan herb",
		Price:       1.99,
		SKU:         "def456", // Changed SKU to be unique
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
		DeletedOn:   time.Time{}, // zero value for time.Time
	},
}
