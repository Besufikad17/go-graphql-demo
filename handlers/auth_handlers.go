package handlers

import (
	"errors"

	models "github.com/Besufikad17/graphqldemo/models"
	services "github.com/Besufikad17/graphqldemo/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewAuthHandler(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) SignUp(user *models.User) (interface{}, error) {
	password := []byte(user.Password)
	var userInDB *models.User

	err := h.DB.Where("email = ? OR phone_number = ?", &user.Email, &user.PhoneNumber).First(&userInDB).Error

	if err == nil {
		return nil, errors.New("Email or Phone number already in use!!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	token, err := services.CreateToken(*user)

	user.Password = string(hashedPassword)
	result := h.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return token, nil
}

func (h handler) Login(loginText string, password string) (interface{}, error) {
	var user *models.User
	var token string

	h.DB.Where("email = ? OR phone_number = ?", loginText, loginText).First(&user)

	if user == nil {
		return nil, errors.New("Invalid credentials!!")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	println(*&user.Email)

	if err == nil {
		newToken, _ := services.CreateToken(*user)
		token = newToken
	} else {
		return nil, errors.New("Invalid credentials!!")
	}
	return token, nil
}
