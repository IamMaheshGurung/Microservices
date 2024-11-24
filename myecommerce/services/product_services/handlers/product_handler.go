package handlers



import(
    "net/http"
    "encoding/json"
    "product_service/models"
    "github.com/gorilla/mux"
)



func GetProducts(w http.ResponseWriter, r * http.Request) {
    products , err:= models.GetAllProducts()
    if err != nil {
        http.Error(w, "product not found", http.StatusNotFound)
        return 
    }

    json.NewEncoder(w).Encode(products)
}



func GetProductByName(w http.ResponseWriter,r* http.Request) {
    name := mux.vars(r) ["name"]

    product, err := models.GetProductByName(name)

    if err != nil {
        http.Error(w, "product not found", http.StatusNotFound)
        return 
    }

    json.NewEncoder(w).Encode(product)

}


func CreateProduct( w http.ResponseWriter, r * http.Request){
    var product models.Product

    json.NewDecoder(r.Body).Decode(product)
    err := models.CreateProduct(product)
    if err != nil {
        http.Error(w, "failed to create the product", http.StatusInternalServerError)
        return 
    }
    w.WriteHeader(http.StatusCreated)
}


func UpdateProduct( w http.ResponseWriter, r * http.Request) {
    name := mux.Vars(r)["name"]

    var product models.Product

    json.NewDecoder(r.Body).Decode(product)
    err := models.UpdateProduct(name, product)
    if err != nil{
        http.Error(w, "Failed to update the product", http.StatusInternalServerError)
        return 
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r * http.Request) {
    name := mux.Vars(r)["name"]
    err := models.DeleteProduct(name) 
    if err != nil {
        http.Error(w, "Failed to delete the product", http.StatusInternalServerError)
        return 
    }
    w.WriteHeader(http.StatusOK)
}
