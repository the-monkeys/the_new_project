package models

import (
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
)

type Article struct {
	Id         string                                   `json:"id"`
	Title      string                                   `json:"title"`
	Content    string                                   `json:"content"`
	Author     string                                   `json:"author"`
	IsDraft    *bool                                    `json:"is_draft"`
	Tags       []string                                 `json:"tags"`
	CreateTime string                                   `json:"create_time"`
	UpdateTime string                                   `json:"update_time"`
	QuickRead  bool                                     `json:"quick_read"`
	CanEdit    *bool                                    `json:"can_edit"`
	OwnerShip  pb.CreateArticleRequest_ContentOwnerShip `json:"content_ownership"`
}

type GetArticleResp struct {
	Author     string `json:"author"`
	CreateTime string `json:"create_time"`
	ID         string `json:"id"`
	QuickRead  string `json:"quick_read"`
	Title      string `json:"title"`
	ViewedBy   string `json:"viewed_by"`
}

//
// type ArticlesForTheMainPage struct {
// 	Shards Shards `json:"_shards"`
// 	Hits   struct {
// 		Hits []struct {
// 			ID     string  `json:"_id"`
// 			Index  string  `json:"_index"`
// 			Score  float64 `json:"_score"`
// 			Source struct {
// 				Author     string `json:"author"`
// 				CreateTime string `json:"create_time"`
// 				ID         string `json:"id"`
// 				QuickRead  string `json:"quick_read"`
// 				Title      string `json:"title"`
// 				ViewedBy   string `json:"viewed_by"`
// 			} `json:"_source"`
// 		} `json:"hits"`
// 		MaxScore float64 `json:"max_score"`
// 		Total    struct {
// 			Relation string `json:"relation"`
// 			Value    int    `json:"value"`
// 		} `json:"total"`
// 	} `json:"hits"`
// 	TimedOut bool `json:"timed_out"`
// 	Took     int  `json:"took"`
// }
type ArticlesForTheMainPage struct {
	Shards struct {
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
				CreateTime string `json:"create_time"`
				QuickRead  bool   `json:"quick_read"`
				Author     string `json:"author"`
				ID         string `json:"id"`
				Title      string `json:"title"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
type Shards struct {
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
}

// END of the Struct

// GetArticleById
type GetArticleById struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Hits []struct {
			ID     string  `json:"_id"`
			Index  string  `json:"_index"`
			Score  float64 `json:"_score"`
			Source struct {
				Author           string `json:"author"`
				CanEdit          string `json:"can_edit"`
				Comments         string `json:"comments"`
				Content          string `json:"content"`
				ContentOwnership string `json:"content_ownership"`
				CreateTime       string `json:"create_time"`
				ID               string `json:"id"`
				IsDraft          string `json:"is_draft"`
				QuickRead        string `json:"quick_read"`
				Tags             string `json:"tags"`
				Title            string `json:"title"`
				UpdateTime       string `json:"update_time"`
				ViewedBy         string `json:"viewed_by"`
			} `json:"_source"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

// End of the struct
