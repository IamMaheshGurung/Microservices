package user



type User struct {
    ID int `json:"id"`
    Username    string `json:"username"`
    PhoneNumber string `json:"phone_number"` 
    Password    string `json:"password"`  
}
