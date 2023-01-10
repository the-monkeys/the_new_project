package models

import "google.golang.org/protobuf/types/known/timestamppb"

type User struct {
	Id         int64
	FirstName  string
	LastName   string
	Email      string
	ProfilePic []byte
	CreateTime *timestamppb.Timestamp
	UpdateTime *timestamppb.Timestamp
	IsActive   bool
	Role       string
	LastLogin  *timestamppb.Timestamp
}
