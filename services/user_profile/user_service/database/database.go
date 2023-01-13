package database

import (
	"database/sql"
	"log"

	"github.com/89minutes/the_new_project/services/user_profile/user_service/models"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserHandler struct {
	GormConn *gorm.DB
	Psql     *sql.DB
}

func Init(url string) UserHandler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalln("cannot connect psql using sql driver, error:, %+v", err)
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return UserHandler{}
	}

	db.AutoMigrate(&models.UserServe{})

	return UserHandler{GormConn: db, Psql: dbPsql}
}
