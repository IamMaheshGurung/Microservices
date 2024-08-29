package data

import (
	"encoding/json"
	"io"
	"log"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decode := json.NewDecoder(r)
	err := decode.Decode(p)
	if err != nil {
		log.Println("Unable to marhsal", err)
	}
	return nil
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:   1,
		Name: "MAhesh Gruung",
		Age:  28,
	},
	{
		ID:   2,
		Name: "Sriana",
		Age:  28,
	},
}
