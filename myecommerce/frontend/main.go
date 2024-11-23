package main

import (
	"html/template"
	"net/http"
	"log"
)

// Product represents a product structure
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	ImageURL    string
}

// Cart represents a cart item structure
type CartItem struct {
	Product  Product
	Quantity int
}

// Global variable to hold products
var products = []Product{
	{ID: 1, Name: "Product 1", Description: "Description of Product 1", Price: 9.99, ImageURL: "/static/images/product1.jpg"},
	{ID: 2, Name: "Product 2", Description: "Description of Product 2", Price: 19.99, ImageURL: "/static/images/product2.jpg"},
}

// Global variable to hold the cart
var cart []CartItem

// RenderHTML renders an HTML template with data
func RenderHTML(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplParsed, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	err = tmplParsed.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

// HomePage displays the homepage with product listing
func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "index.html", products)
}

// ProductPage displays the product details page
func ProductPage(w http.ResponseWriter, r *http.Request) {
	// Hardcoding product ID for simplicity
	product := products[0]
	RenderHTML(w, "product.html", product)
}

// CartPage displays the shopping cart
func CartPage(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "cart.html", cart)
}

// AddToCart adds a product to the cart
func AddToCart(w http.ResponseWriter, r *http.Request) {
	// For simplicity, just add Product 1 to the cart each time
	product := products[0]
	cart = append(cart, CartItem{Product: product, Quantity: 1})
	RenderHTML(w, "cart.html", cart)
}

// CheckoutPage displays the checkout page
func CheckoutPage(w http.ResponseWriter, r *http.Request) {
	RenderHTML(w, "checkout.html", cart)
}

// main function to start the server
func main() {
	// Serve static files like images, CSS, etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Route handlers
	http.HandleFunc("/", HomePage)          // Home page route
	http.HandleFunc("/product", ProductPage) // Product details route
	http.HandleFunc("/cart", CartPage)      // Cart route
	http.HandleFunc("/cart/add", AddToCart) // Add to cart route
	http.HandleFunc("/checkout", CheckoutPage) // Checkout route

	// Start the server
	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

