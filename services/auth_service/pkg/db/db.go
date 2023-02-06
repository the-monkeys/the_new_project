package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type AuthDBHandler struct {
	PsqlClient *sql.DB
}

//
func NewAuthDBHandler(url string) (*AuthDBHandler, error) {
	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
		return nil, err
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return nil, err

	}

	return &AuthDBHandler{PsqlClient: dbPsql}, nil
}
