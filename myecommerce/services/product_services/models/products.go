package models



import(
    "errors"
    "gorm.io/gorm"
)




type Product struct {
    gorm.Model
    Name string `json:"name"`
    Price int `json:"price"`
    Stock int `json:"stock"`
}


var db *gorm.DB
var products []Product
var product Product




func InitializeDB(database * gorm.DB){
    db = database
    db.AutoMigrate(&Product{})
}


func GetAllProducts() ([]Product, error) {
    result := db.Find(&products)
    return products, result.Error
   
}







func GetProductByName(name string) (*Product, error) {
    result := db.First(&product, "name=?", name)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, errors.New("product not found")
        }
    
    return &product, result.Error
}



func CreateProduct(product Product) error {
    result := db.Create(&product)
    return result.Error
}


func UpdateProduct(name string, updatedProduct Product) error {
    result := db.First(&product, "name=?", name)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return errors.New("product not found")
        }
    
        product.Name = updatedProduct.Name
        product.Price = updatedProduct.Price
        product.Stock = updatedProduct.Stock
        return db.Save(&product).Error
    }



