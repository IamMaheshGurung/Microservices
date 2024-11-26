package handler




import(
    "net/http"
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


func (s *UserService) CreateUser(username, phonenumber, password string) (*models.User, error){

    var existingUser models.User


    if err := s.db.Where("phone_number =?", phonenumber).First(&existingUser).Error; err != nil {
        return nil, errors.New("Username or mobile has been already register, try direct logging")
    }


    if err := s.db.Where("username =?", username).First(&existingUser).Error; err != nil {
        return nil, errors.New("Username or mobile has been already register, try another username")
    }


    

    user := models.User {
        Username: username,
        PhoneNumber: phonenumber,
        Password: password,
    }


    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil 
}



