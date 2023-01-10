package database

import (
	"log"

	"github.com/89minutes/the_new_project/auth_service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func Init(url string) UserHandler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return UserHandler{db}
}
