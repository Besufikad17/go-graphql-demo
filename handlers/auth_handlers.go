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
	var existingUser models.User
	err := h.DB.Where("email = ? OR phone_number = ?", &user.Email, &user.PhoneNumber).Find(&existingUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		password := []byte(user.Password)
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
	} else {
		return nil, errors.New("Email or phone number already in use")
	}
}

func (h handler) Login(loginText string, password string) (interface{}, error) {
	var user *models.User
	var token string

	h.DB.Where("email = ? OR phone_number = ?", loginText, loginText).First(&user)

	if user == nil {
		return nil, errors.New("Invalid credentials!!")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err == nil {
		newToken, _ := services.CreateToken(*user)
		token = newToken
	} else {
		return nil, errors.New("Invalid credentials!!")
	}
	return token, nil
}
