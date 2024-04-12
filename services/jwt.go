package services

import (
	models "github.com/Besufikad17/graphqldemo/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func CreateToken(user models.User) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
			"exp":         time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken() {

}
