package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/database"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/utils"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
)

type openSearchClient struct {
	client *opensearch.Client
	log    *logrus.Logger
}

func newOpenSearchClient(url, username, password string, log *logrus.Logger) (*openSearchClient, error) {
	client, err := database.NewOSClient(url, username, password)
	if err != nil {
		logrus.Errorf("Failed to connect to opensearch instance, error: %+v", err)
		return nil, err
	}

	return &openSearchClient{
		client: client,
		log:    log,
	}, nil
}

// TODO: Implement MapArticleIndex function
// MapArticleIndex setups the mapping considering database schema and search factors.
func (oso *openSearchClient) MapArticleIndex(mapping string) error {
	return nil
}

// CreateAnArticle creates a document for an article posted by a user
func (oso *openSearchClient) CreateAnArticle(article models.Article) (*opensearchapi.Response, error) {
	oso.log.Infof("received an article with id: %s", article.Id)

	bs, err := json.Marshal(article)
	if err != nil {
		oso.log.Errorf("cannot marshal the article, error: %v", err)
		return nil, err
	}

	document := strings.NewReader(string(bs))

	osReq := opensearchapi.IndexRequest{
		Index:      utils.OpensearchArticleIndex,
		DocumentID: article.Id,
		Body:       document,
	}

	insertResponse, err := osReq.Do(context.Background(), oso.client)
	if err != nil {
		oso.log.Errorf("error while creating/drafting article, error: %+v", err)
		return insertResponse, err
	}

	if insertResponse.IsError() {
		oso.log.Errorf("error creating an article, insert response: %+v", insertResponse)
		return insertResponse, err
	}

	oso.log.Infof("successfully created an article for user: %s, insert response: %+v", article.AuthorEmail, insertResponse)
	return insertResponse, nil
}
