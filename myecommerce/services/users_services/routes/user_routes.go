package routes




import(
    "userService/handlers"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
    
)



func SetupRoutes(db *gorm.DB) * mux.Router {
    r := mux.NewRouter()
    userService := handlers.NewUserService(db)


    r.HandleFunc("/users", userService.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", userService.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", userService.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", userService.DeleteUser).Methods("DELETE")
return r
}
