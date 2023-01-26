package db

import (
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/sirupsen/logrus"
)

func (auth *AuthDBHandler) RegisterUser(user models.TheMonkeysUser) error {
	stmt, err := auth.PsqlClient.Prepare(`INSERT INTO the_monkeys_user (
		unique_id, first_name, last_name, email, password, create_time, update_time, is_active, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`)

	if err != nil {
		logrus.Errorf("cannot prepare statement to register user for %s error: %+v", user.Email, err)
		return err
	}

	result, err := stmt.Exec(user.UUID, user.FirstName, user.LastName, user.Email,
		user.Password, user.CreateTime, user.UpdateTime, user.IsActive, user.Role)

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
