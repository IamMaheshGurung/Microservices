
package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres" 
    "fmt"
)

var DB *gorm.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() (*gorm.DB, error) {
    // PostgreSQL connection string
    connstr := "user=postgres dbname=mydb password=gurung67 host=localhost port=4000 sslmode=disable"
    
    // Open the database connection
    db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Check if the connection is alive by pinging the database
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    err = sqlDB.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Database connection established successfully")
    return db, nil
}

