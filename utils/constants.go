package graphqldemo

import (
	models "github.com/Besufikad17/graphqldemo/models"
)

var Users = []models.User{
	{
		ID:          1,
		FirstName:   "Abebe",
		LastName:    "Kebede",
		Email:       "abe@gmail.com",
		PhoneNumber: "0912345678",
	},
	{
		ID:          2,
		FirstName:   "Kebede",
		LastName:    "Abebe",
		Email:       "kebe@gmail.com",
		PhoneNumber: "0978564321",
	},
}
