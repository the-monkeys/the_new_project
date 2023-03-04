package utils

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Errors(err error) error {
	switch err {
	case sql.ErrNoRows:
		logrus.Infof("cannot find the row")
		return status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
	case sql.ErrTxDone:
		logrus.Infof("The transaction has already been committed or rolled back.")
		return status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
	case sql.ErrConnDone:
		logrus.Infof("The database connection has been closed.")
		return status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
	default:
		logrus.Infof("An internal server error occurred: %v\n", err)
		return status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
	}
}
