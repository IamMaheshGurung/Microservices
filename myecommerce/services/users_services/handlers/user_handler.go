package handler




import(
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "errors"
    "gorm.io/gorm"
    "userService/models"
    "log"
    

)



type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}


func (s *UserService) CreateUser(username, phonenumber, password string) (*models.User, error){

    var existingUser models.User


    if err := s.db.Where("phone_number =?", phonenumber).First(&existingUser).Error; err != nil {
        return nil, errors.New("Username or mobile has been already register, try direct logging")
    }


    if err := s.db.Where("username =?", username).First(&existingUser).Error; err != nil {
        return nil, errors.New("Username or mobile has been already register, try another username")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
           log.Printf("Unable to get the hashed paswword")
            
       }
       


    newPassword := string(hashedPassword)
    

    user := models.User {
        Username: username,
        PhoneNumber: phonenumber,
        Password: newPassword,
    }


    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil 
}

func (s *UserService) GetUser(id int) (*models.User, error) {
    var user models.User 
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("User with ID %d not found", id)
    }
    return &user, nil 
}


func (s *UserService) UpdateUser(id int, username, phoneNumber, password string) (*models.User, error) {
    var user models.User 


    if err := s.db.First(&user, id).Error; err != nil {
        return nil, fmt.Errorf("user id %d not found or some error", id)
    }

    var existingUser models.User

    
    if err := s.db.Where("phone_number = ? AND id != ?", phoneNumber, id).First(&existingUser).Error; err != nil {
        return nil, errors.New("Phone number already in use")
    }

     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
           log.Printf("Unable to get the hashed paswword")
            
       }
       


    newPassword := string(hashedPassword)
    

            user.Username = username
            user.PhoneNumber = phoneNumber
            user.Password = newPassword



            if err := s.db.Save(&user).Error; err != nil {
                return nil, err 
            }
            return &user, nil 
}


func(s *UserService) DeleteUser(id int) error {
    if err := s.db.Delete(&models.User{}, id).Error; err != nil {
        return fmt.Errorf("user with ID %d not found", id)
    }
    return nil 
}


