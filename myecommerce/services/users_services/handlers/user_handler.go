
package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "gorm.io/gorm"
    "userService/models"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)
type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}


func (s *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
    var newUser models.User

    // Parse the incoming request JSON into the newUser struct
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Validate phone number and username
    var existingUser models.User

    if err := s.db.Where("phone_number = ?", newUser.PhoneNumber).First(&existingUser).Error; err == nil {
        http.Error(w, "Phone number is already registered", http.StatusConflict)
        return
    }

    if err := s.db.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
        http.Error(w, "Username is already taken, try another one", http.StatusConflict)
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Unable to hash the password: %v", err)
        http.Error(w, "Unable to process password", http.StatusInternalServerError)
        return
    }
    newUser.Password = string(hashedPassword)

    // Save the new user to the database
    if err := s.db.Create(&newUser).Error; err != nil {
        http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
        return
    }

    // Respond with the created user as a JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}


//for getting user

func (s *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"] 
    var user models.User

    if err := s.db.First(&user, id).Error; err != nil {
        http.Error(w, fmt.Sprintf("User with ID %s not found", id), http.StatusNotFound)
        return
    }

    // Respond with the user in JSON format
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

//for updating the user

func (s *UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var user models.User

    if err := s.db.First(&user, id).Error; err != nil {
        http.Error(w, fmt.Sprintf("User with ID %s not found", id), http.StatusNotFound)
        return
    }

    // Update user data
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Save the updated user
    if err := s.db.Save(&user).Error; err != nil {
        http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
        return
    }

    // Respond with the updated user
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}


func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    if err := s.db.Delete(&models.User{}, id).Error; err != nil {
        http.Error(w, fmt.Sprintf("Failed to delete user with ID %s: %v", id, err), http.StatusInternalServerError)
        return
    }

    // Respond with a success message
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

