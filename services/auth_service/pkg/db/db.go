package db

import (
	"log"

	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		logrus.Errorf("cannot migrate the user table, error: %v", err)
	}
	if err = db.AutoMigrate(&models.PasswordReset{}); err != nil {
		logrus.Errorf("cannot migrate the pass reset table, error: %v", err)
	}

	return Handler{db}
}
