package db

import (
	"database/sql"

	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	GormConn *gorm.DB
	Psql     *sql.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("cannot open through gorm driver, error:, %+v", err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		logrus.Errorf("cannot migrate the user table, error: %v", err)
	}
	if err = db.AutoMigrate(&models.PasswordReset{}); err != nil {
		logrus.Errorf("cannot migrate the pass reset table, error: %v", err)
	}

	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
		return Handler{}
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return Handler{}
	}

	return Handler{GormConn: db, Psql: dbPsql}
}
