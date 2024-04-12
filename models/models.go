package graphqldemo

import "gorm.io/gorm"

type Role string

const (
	Admin    Role = "ADMIN"
	Customer Role = "CUSTOMER"
)

type User struct {
	gorm.Model
	ID          uint   `gorm: "primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        Role   `gorm:"type:user_role"`
}

type Message struct {
	Text string
}

type AuthResponse struct {
	Message string
	Token   string
}
