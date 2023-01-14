package models

import "time"

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic []byte `json:"profilePic,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	IsActive   bool   `json:"isActive,omitempty"`
	Role       int32  `json:"role,omitempty"`
	LastLogin  string `json:"lastLogin,omitempty"`
}

type PasswordReset struct {
	Id                int64     `json:"id" gorm:"primaryKey"`
	Email             string    `json:"email"`
	RecoveryHash      string    `json:"recovery_hash"`
	TimeOut           time.Time `json:"time_out"`
	LastPasswordReset time.Time `json:"last_pass_reset"`
}
