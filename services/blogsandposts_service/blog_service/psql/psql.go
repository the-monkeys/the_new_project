package psql

import (
	"database/sql"

	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/models"
	"github.com/sirupsen/logrus"
)

type PostDBHandler struct {
	psqlClient *sql.DB
	logger     *logrus.Logger
}

func NewPostDBhandler(client *sql.DB, logger *logrus.Logger) *PostDBHandler {
	return &PostDBHandler{psqlClient: client, logger: logger}
}

func (psql *PostDBHandler) StoreAPost(post models.Blogs) error {
	stmt, err := psql.psqlClient.Prepare(`INSERT INTO the_monkeys_post (
		uuid, title, html_content, raw_content, author_name, author_id, published, tags, create_time, update_time, can_edit, content_ownership, folder_path)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`)

	if err != nil {
		logrus.Errorf("cannot prepare statement to register user for %s error: %+v", post.Title, err)
		return err
	}

	result, err := stmt.Exec(post.Id, post.Title, post.ContentFormatted, post.ContentRaw,
		post.AuthorName, post.AuthorId, post.Published, post.Tags, post.CreateTime, post.UpdateTime, post.CanEdit, post.OwnerShip, post.FolderPath)

	if err != nil {
		logrus.Errorf("cannot execute register user query for %s, error: %v", post.Title, err)
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %s, error: %v", post.Title, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", post.Title, err)
		return err
	}

	return nil
}
