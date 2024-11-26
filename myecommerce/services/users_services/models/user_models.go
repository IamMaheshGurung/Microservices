package models


import(
    "gorm.io/gorm"
)

type User struct {
    gorm.Model 
    Username    string `gorm:"unique;not null" json:"username"`
    PhoneNumber string `gorm:"unique;not null" json:"phone_number"`
    Password    string `gorm:"not null" json:"password"` 
}


