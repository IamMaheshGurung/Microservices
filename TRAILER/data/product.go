package data

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{ID: 1,
		Name: "MAhesh Gruung",
		Age:  28,
	},
}
