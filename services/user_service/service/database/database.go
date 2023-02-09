package database

import (
	"database/sql"
	"errors"

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
		Instagram:     profile.Instagram,
		Twitter:       profile.Twitter,
		EmailVerified: profile.EmailVerified,
	}
	return res, nil
}

func (uh *UserDbHandler) UpdateMyProfile(info *pb.SetMyProfileReq) error {
	stmt, err := uh.Psql.Prepare(`UPDATE the_monkeys_user SET first_name=$1, last_name=$2,
	country_code=$3, mobile_no=$4, about=$5, instagram=$6, twitter=$7 WHERE email=$8`)
	if err != nil {
		uh.log.Errorf("cannot prepare update profile statement, error: %v", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(info.FirstName, info.LastName, info.CountryCode, info.MobileNo,
		info.About, info.Instagram, info.Twitter, info.Email)
	if err != nil {
		uh.log.Errorf("cannot execute update profile statement, error: %v", err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", info.Email, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", info.Email, err)
		return errors.New("more or less than 1 row is affected")
	}

	return nil
}

func (uh *UserDbHandler) UploadProfilePic(pic []byte, id int64) error {
	stmt, err := uh.Psql.Prepare(`UPDATE the_monkeys_user SET profile_pic=$1 WHERE id=$2`)
	if err != nil {
		uh.log.Errorf("cannot prepare upload profile pic statement, error: %v", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(pic, id)
	if err != nil {
		uh.log.Errorf("cannot execute update profile pic statement, error: %v", err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", id, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", id, err)
		return errors.New("more or less than 1 row is affected")
	}

	return nil
}
