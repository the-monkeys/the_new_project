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
