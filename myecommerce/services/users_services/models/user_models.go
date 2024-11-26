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

package user

import (
    "errors"
    "fmt"
    "gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
    ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Username    string `gorm:"unique;not null" json:"username"`
    PhoneNumber string `gorm:"unique;not null" json:"phone_number"`
    Password    string `gorm:"not null" json:"password"` // In a real app, you should hash the password
}

// UserService defines methods for handling user operations.
type UserService struct {
    db *gorm.DB
}

// NewUserService creates a new instance of UserService.
func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

// CreateUser creates a new user and returns the user.
func (s *UserService) CreateUser(username, phoneNumber, password string) (*User, error) {
    // Check if phone number already exists
    var existingUser User
    if err := s.db.Where("phone_number = ?", phoneNumber).First(&existingUser).Error; err == nil {
        return nil, errors.New("phone number already in use")
    }

    // Check if username already exists
    if err := s.db.Where("username = ?", username).First(&existingUser).Error; err == nil {
        return nil, errors.New("username already taken")
    }

    // Create a new user
    user := User{
        Username:    username,
        PhoneNumber: phoneNumber,
        Password:    password, // In a real app, you should hash the password
    }

    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(id int) (*User, error) {
    var user User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }
    return &user, nil
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(id int, username, phoneNumber, password string) (*User, error) {
    var user User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }

    // Ensure the new phone number is unique
    var existingUser User
    if err := s.db.Where("phone_number = ? AND id != ?", phoneNumber, id).First(&existingUser).Error; err == nil {
        return nil, errors.New("phone number already in use")
    }

    // Update the user fields
    user.Username = username
    user.PhoneNumber = phoneNumber
    user.Password = password // In a real app, you should hash the password

    if err := s.db.Save(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id int) error {
    if err := s.db.Delete(&User{}, id).Error; err != nil {
        return fmt.Errorf("user with ID %d not found", id)
    }
    return nil
}

