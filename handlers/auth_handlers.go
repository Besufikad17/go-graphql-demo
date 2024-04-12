package handlers

import (
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

func Login(loginText string, password string) (interface{}, error) {
	return nil, nil
}
