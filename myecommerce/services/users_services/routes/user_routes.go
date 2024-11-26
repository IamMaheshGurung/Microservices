package routes




import(
    "userService/handlers"
    "github.com/gorilla/mux"
)



func SetupRoutes() * mux.Router {
    router := mux.NewRouter()



    router.HandleFunc("/login", handler.NewUserService.
