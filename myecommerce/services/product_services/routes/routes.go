package routes



import(
    
    "product_service/handlers"
    "github.com/gorilla/mux"
)




func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
    router.HandlerFunc("/products/{name}", handlers.GetProductByName).Methods("GET")
    router.HandlerFunc("/products", handlers.CreateProduct).Methods("POST")
    router.HandlerFunc("/products/{name}", handlers.UpdateProduct).Methods("PUT")
    router.HandlerFunc("/products/{name}", handlers.DeleteProduct).Methods("DELETE")

    return router

}

