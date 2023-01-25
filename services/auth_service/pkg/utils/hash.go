package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("bcrypt cannot compare hash, error %+v", err)
		return ""
	}
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logrus.Errorf("bcrypt cannot compare hash, error %+v", err)
		return false
	}
	return true
}
