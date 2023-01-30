package db

import (
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/sirupsen/logrus"
)

func (auth *AuthDBHandler) RegisterUser(user models.TheMonkeysUser) error {
	stmt, err := auth.PsqlClient.Prepare(`INSERT INTO the_monkeys_user (
		unique_id, first_name, last_name, email, password, create_time, 
		update_time, is_active, role, email_verification_token, 
		email_verification_timeout) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`)
	defer stmt.Close()
	if err != nil {
		logrus.Errorf("cannot prepare statement to register user for %s error: %+v", user.Email, err)
		return err
	}

	result, err := stmt.Exec(user.UUID, user.FirstName, user.LastName, user.Email,
		user.Password, user.CreateTime, user.UpdateTime, user.IsActive, user.Role, user.EmailVerificationToken, user.EmailVerificationTimeout)

	if err != nil {
		logrus.Errorf("cannot execute register user query for %s, error: %v", user.Email, err)
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", user.Email, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", user.Email, err)
		return err
	}

	return nil
}

func (auth *AuthDBHandler) UpdateEmailVerToken(user models.TheMonkeysUser) error {
	stmt, err := auth.PsqlClient.Prepare(`UPDATE the_monkeys_user SET 
	email_verification_token=$1, 
	email_verification_timeout=$2 WHERE email=$3;`)
	defer stmt.Close()
	if err != nil {
		logrus.Errorf("cannot prepare statement to update verify email token for %s error: %+v", user.Email, err)
		return err
	}

	res, err := stmt.Exec(user.EmailVerificationToken, user.EmailVerificationTimeout, user.Email)
	if err != nil {
		logrus.Errorf("cannot update the verification details for %s, error: %v", user.Email, err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", user.Email, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", user.Email, err)
		return err
	}
	return nil
}
