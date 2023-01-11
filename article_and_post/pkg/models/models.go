package models

import "google.golang.org/protobuf/types/known/timestamppb"

type Article struct {
	Id         string                 `json:"id"`
	Title      string                 `json:"title"`
	Content    string                 `json:"content"`
	Author     string                 `json:"author"`
	IsDraft    bool                   `json:"is_draft"`
	Tags       []string               `json:"tags"`
	CreateTime *timestamppb.Timestamp `json:"create_time"`
	UpdateTime *timestamppb.Timestamp `json:"update_time"`
	QuickRead  bool                   `json:"quick_read"`
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
type ArticlesForTheMainPage struct {
	Shards Shards `json:"_shards"`
	Hits   struct {
		Hits []struct {
			ID     string  `json:"_id"`
			Index  string  `json:"_index"`
			Score  float64 `json:"_score"`
			Source struct {
				Author     string `json:"author"`
				CreateTime string `json:"create_time"`
				ID         string `json:"id"`
				QuickRead  string `json:"quick_read"`
				Title      string `json:"title"`
				ViewedBy   string `json:"viewed_by"`
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
type Shards struct {
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
}

// END of the Struct
