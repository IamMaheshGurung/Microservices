package routes



import(
    
    "product_service/handlers"
    "github.com/gorilla/mux"
)




func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
    router.HandleFunc("/products/{name}", handlers.GetProductByName).Methods("GET")
    router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
    router.HandleFunc("/products/{name}", handlers.UpdateProduct).Methods("PUT")
    router.HandleFunc("/products/{name}", handlers.DeleteProduct).Methods("DELETE")

    return router

}

