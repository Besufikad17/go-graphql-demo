package db

import (
	models "github.com/Besufikad17/graphqldemo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Init() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DB_CONNECTION")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Print("Database connected successfully")
	}

	log.Print("Migration started")
	db.AutoMigrate(&models.User{})
	log.Print("Migration ended")

	return db
}
