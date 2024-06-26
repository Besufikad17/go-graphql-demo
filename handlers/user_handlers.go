package handlers

import (
	"errors"
	models "github.com/Besufikad17/graphqldemo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) AddUser(user *models.User) (interface{}, error) {
	var existingUser models.User
	err := h.DB.Where("email = ? OR phone_number = ?", &user.Email, &user.PhoneNumber).Find(&existingUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		password := []byte(user.Password)

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
		result := h.DB.Create(user)
		if result.Error != nil {
			return nil, result.Error
		}
		return user, nil
	} else {
		return nil, errors.New("Email or phone number already in use")
	}
}

func (h handler) GetAllUsers(skip *int, take *int, text *string) (interface{}, error) {
	var users []models.User
	h.DB.Limit(*take).Offset(*skip).Where("first_name LIKE ? OR last_name LIKE ?", "%"+*text+"%", "%"+*text+"%").Find(&users)
	return users, nil
}

func (h handler) GetUserById(id int) (interface{}, error) {
	var user = models.User{ID: uint(id)}
	h.DB.Find(&user)
	return user, nil
}

func (h handler) DeleteUser(id int) (interface{}, error) {
	var user = models.User{ID: uint(id)}
	h.DB.Delete(&user)
	return nil, nil
}

func (h handler) UpdateUser(id uint, user *models.User) (interface{}, error) {
	h.DB.Model(user).Where("id = ?", id).Updates(user)
	return user, nil
}
