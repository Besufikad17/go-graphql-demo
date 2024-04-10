package handlers

import (
	models "github.com/Besufikad17/graphqldemo/models"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) AddUser(user *models.User) (interface{}, error) {
	result := h.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (h handler) GetAllUsers(skip *int, take *int, text *string) (interface{}, error) {
	var users []models.User
	h.DB.Limit(*take).Offset(*skip).Where("first_name LIKE ? OR last_name LIKE ?", "%"+*text+"%", "%"+*text+"%").Find(&users)
	return users, nil
}
