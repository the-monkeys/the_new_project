package models

import (
	"time"

	"github.com/89minutes/the_new_project/services/auth_service/pkg/pb"
)

type TheMonkeysUser struct {
	Id                        int64                          `json:"id" gorm:"primaryKey"`
	UUID                      string                         `json:"unique_id"`
	FirstName                 string                         `json:"first_name"`
	LastName                  string                         `json:"last_name"`
	Email                     string                         `json:"email"`
	Password                  string                         `json:"password"`
	CreateTime                string                         `json:"create_time,omitempty"`
	UpdateTime                string                         `json:"update_time,omitempty"`
	IsActive                  bool                           `json:"is_active,omitempty"`
	Role                      int32                          `json:"role,omitempty"`
	EmailVerificationToken    string                         `json:"email_verification_token"`
	EmailVerificationTimeout  time.Time                      `json:"email_verification_timeout"`
	MobileVerificationToken   string                         `json:"mobile_verification_token"`
	MobileVerificationTimeout time.Time                      `json:"mobile_verification_timeout"`
	Deactivated               bool                           `json:"deactivated"`
	LoginMethod               pb.RegisterRequest_LoginMethod `json:"login_method"`
}

type PasswordReset struct {
	Id                int64     `json:"id" gorm:"primaryKey"`
	Email             string    `json:"email"`
	RecoveryHash      string    `json:"recovery_hash"`
	TimeOut           time.Time `json:"time_out"`
	LastPasswordReset time.Time `json:"last_pass_reset"`
}
