package database

import (
	"crypto/tls"
	"net/http"

	"github.com/opensearch-project/opensearch-go"
)

func NewOSClient(url, username, password string) (*opensearch.Client, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{url},
		Username:  username, // For testing only. Don't store credentials in code.
		Password:  password,
	})

	return client, err
}
