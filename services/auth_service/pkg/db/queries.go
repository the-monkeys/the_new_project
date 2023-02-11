package db

import (
	"database/sql"
	"time"

	"github.com/89minutes/the_new_project/common"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/models"
	"github.com/89minutes/the_new_project/services/auth_service/pkg/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (auth *AuthDBHandler) GetNamesEmailFromEmail(req *pb.ForgotPasswordReq) (*models.TheMonkeysUser, error) {
	var user models.TheMonkeysUser

	if err := auth.PsqlClient.QueryRow("SELECT id, first_name, last_name, email from the_monkeys_user where email=$1;", req.GetEmail()).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email); err != nil {
		switch err {
		case sql.ErrNoRows:
			logrus.Errorf("cannot fine the user with id %v ", req.GetEmail())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			logrus.Errorf("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			logrus.Errorf("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			logrus.Errorf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}

	return &user, nil
}

func (auth *AuthDBHandler) UpdatePasswordRecoveryToken(hash string, req *models.TheMonkeysUser) error {
	// TODO: start a database transaction from here till all the process are complete
	sqlStmt, err := auth.PsqlClient.Prepare(`INSERT INTO pw_reset (
		user_id, email, recovery_hash, time_out, last_password_reset) 
		VALUES ($1, $2, $3, $4, $5);`)
	if err != nil {
		logrus.Errorf("cannot prepare the reset link for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	result, err := sqlStmt.Exec(req.Id, req.Email, hash, time.Now().Add(time.Minute*5), time.Now().Format(common.DATE_TIME_FORMAT))
	if err != nil {
		logrus.Errorf("cannot sent the reset link for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("cannot check for affected rows for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
	if affectedRows != 1 {
		logrus.Errorf("more than 1 rows are getting affected for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	return nil
}

func (auth *AuthDBHandler) UpdatePassword(password, email string) error {
	stmt, err := auth.PsqlClient.Prepare(`UPDATE the_monkeys_user SET 
	password=$1 WHERE email=$2;`)
	defer stmt.Close()
	if err != nil {
		logrus.Errorf("cannot prepare statement to update password for %s error: %+v", email, err)
		return err
	}

	res, err := stmt.Exec(password, email)
	if err != nil {
		logrus.Errorf("cannot update the password for %s, error: %v", email, err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", email, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", email, err)
		return err
	}
	return nil
}
