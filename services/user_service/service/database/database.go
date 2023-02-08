package database

import (
	"database/sql"

	"github.com/89minutes/the_new_project/services/user_service/service/models"
	"github.com/89minutes/the_new_project/services/user_service/service/pb"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type UserDbHandler struct {
	Psql *sql.DB
	log  *logrus.Logger
}

func NewUserDbHandler(url string, log *logrus.Logger) *UserDbHandler {
	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return nil
	}

	return &UserDbHandler{Psql: dbPsql, log: log}
}

func (uh *UserDbHandler) GetMyProfile(id int64) (*pb.GetMyProfileRes, error) {
	profile := &models.MyProfile{}
	countryCode := sql.NullString{}
	if err := uh.Psql.QueryRow(`SELECT id, first_name, last_name, email, create_time,
	is_active, country_code, mobile_no, about, instagram, twitter, email_verified FROM
	the_monkeys_user WHERE id=$1`, id).Scan(&profile.Id, &profile.FirstName, &profile.LastName,
		&profile.Email, &profile.CreateTime, &profile.IsActive, &countryCode, &profile.Mobile,
		&profile.About, &profile.Instagram, &profile.Twitter, &profile.EmailVerified); err != nil {
		return nil, err
	}

	res := &pb.GetMyProfileRes{
		Id:            profile.Id,
		FirstName:     profile.FirstName,
		LastName:      profile.LastName,
		Email:         profile.Email,
		CreateTime:    nil,
		IsActive:      profile.IsActive,
		CountryCode:   profile.CountryCode,
		Mobile:        profile.Mobile,
		About:         profile.About,
		Insagram:      profile.Instagram,
		Twitter:       profile.Twitter,
		EmailVerified: profile.EmailVerified,
	}
	return res, nil
}
