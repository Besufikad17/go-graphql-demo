package graphqldemo

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm: "primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
