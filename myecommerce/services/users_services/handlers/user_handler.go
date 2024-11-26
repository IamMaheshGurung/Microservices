
package handlers

import (
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
    "errors"
    "gorm.io/gorm"
    "userService/models"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

// CreateUser creates a new user and returns the user object.
func (s *UserService) CreateUser(username, phonenumber, password string) (*models.User, error) {
    var existingUser models.User

    // Check if phone number already exists
    if err := s.db.Where("phone_number = ?", phonenumber).First(&existingUser).Error; err == nil {
        return nil, errors.New("phone number is already registered")
    }

    // Check if username already exists
    if err := s.db.Where("username = ?", username).First(&existingUser).Error; err == nil {
        return nil, errors.New("username is already taken, try another one")
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Unable to hash the password: %v", err)
        return nil, errors.New("unable to process password")
    }

    // Create the new user object
    user := models.User{
        Username:    username,
        PhoneNumber: phonenumber,
        Password:    string(hashedPassword), // Store the hashed password
    }

    // Save the new user to the database
    if err := s.db.Create(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }

    return &user, nil
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(id int) (*models.User, error) {
    var user models.User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }
    return &user, nil
}

// UpdateUser updates an existing user's information.
func (s *UserService) UpdateUser(id int, username, phoneNumber, password string) (*models.User, error) {
    var user models.User

    // Check if the user exists
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }

    // Ensure the new phone number is unique
    var existingUser models.User
    if err := s.db.Where("phone_number = ? AND id != ?", phoneNumber, id).First(&existingUser).Error; err == nil {
        return nil, errors.New("phone number already in use")
    }

    // Hash the password if it's provided
    if password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            log.Printf("Unable to hash the password: %v", err)
            return nil, errors.New("unable to process password")
        }
        user.Password = string(hashedPassword) // Update the password
    }

    // Update the other fields
    user.Username = username
    user.PhoneNumber = phoneNumber

    // Save the updated user to the database
    if err := s.db.Save(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to update user: %v", err)
    }

    return &user, nil
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id int) error {
    // Try to delete the user
    if err := s.db.Delete(&models.User{}, id).Error; err != nil {
        return fmt.Errorf("failed to delete user with ID %d: %v", id, err)
    }
    return nil
}

