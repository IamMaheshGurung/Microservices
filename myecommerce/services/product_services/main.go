package main 



import(
    "fmt"
    "log"
    "net/http"

)



func main(){
    router := routes.SetupRoutes()
    log.Println("Product service running at the port 8081")
    log.Fatal(http.ListenAndServe(":8081", router))

}

