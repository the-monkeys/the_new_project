package models

import (
	"time"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
)

type Article struct {
	Id          string                                   `json:"id"`
	Title       string                                   `json:"title"`
	Content     string                                   `json:"content"`
	Author      string                                   `json:"author"`
	AuthorEmail string                                   `json:"author_email"`
	IsDraft     *bool                                    `json:"is_draft"`
	Tags        []string                                 `json:"tags"`
	CreateTime  string                                   `json:"create_time"`
	UpdateTime  string                                   `json:"update_time"`
	QuickRead   bool                                     `json:"quick_read"`
	CanEdit     *bool                                    `json:"can_edit"`
	OwnerShip   pb.CreateArticleRequest_ContentOwnerShip `json:"content_ownership"`
	FolderPath  string                                   `json:"folder_path"`
}

type GetArticleResp struct {
	Author     string `json:"author"`
	CreateTime string `json:"create_time"`
	ID         string `json:"id"`
	QuickRead  string `json:"quick_read"`
	Title      string `json:"title"`
	ViewedBy   string `json:"viewed_by"`
}

type Last100Articles struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore interface{} `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			ID     string      `json:"_id"`
			Score  interface{} `json:"_score"`
			Source struct {
				CreateTime  time.Time `json:"create_time"`
				QuickRead   bool      `json:"quick_read"`
				Author      string    `json:"author"`
				AuthorEmail string    `json:"author_email"`
				ID          string    `json:"id"`
				Title       string    `json:"title"`
			} `json:"_source"`
			Sort []int64 `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
}

// END of the Struct

// GetArticleById
type GetArticleById struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ID               string   `json:"id"`
				Title            string   `json:"title"`
				Content          string   `json:"content"`
				Author           string   `json:"author"`
				IsDraft          bool     `json:"is_draft"`
				Tags             []string `json:"tags"`
				CreateTime       string   `json:"create_time"`
				UpdateTime       string   `json:"update_time"`
				QuickRead        bool     `json:"quick_read"`
				CanEdit          bool     `json:"can_edit"`
				ContentOwnership int      `json:"content_ownership"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// End of the struct
