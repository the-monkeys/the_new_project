package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthDBHandler struct {
	GormClient *gorm.DB
	PsqlClient *sql.DB
}

// func Init(url string) AuthDBHandler {
// 	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
// 	if err != nil {
// 		logrus.Fatalf("cannot open through gorm driver, error:, %+v", err)
// 	}

// 	dbPsql, err := sql.Open("postgres", url)
// 	if err != nil {
// 		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
// 		return AuthDBHandler{}
// 	}

// 	if err = dbPsql.Ping(); err != nil {
// 		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
// 		return AuthDBHandler{}
// 	}

// 	return AuthDBHandler{GormConn: db, Psql: dbPsql}
// }

//
func NewAuthDBHandler(url string) (*AuthDBHandler, error) {
	dbGorm, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("cannot open through gorm driver, error:, %+v", err)
	}

	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
		return nil, err
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return nil, err

	}

	return &AuthDBHandler{GormClient: dbGorm, PsqlClient: dbPsql}, nil
}
