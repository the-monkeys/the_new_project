package service

import (
	"log"
	"net/http"

	"github.com/89minutes/the_new_project/services/article_and_post/pkg/models"
	"github.com/89minutes/the_new_project/services/article_and_post/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ParseToStruct(result models.Last100Articles) []pb.GetArticlesResponse {
	var resp []pb.GetArticlesResponse

	for _, val := range result.Hits.Hits {

		res := pb.GetArticlesResponse{
			Id:          val.Source.ID,
			Title:       val.Source.Title,
			Author:      val.Source.Author,
			AuthorEmail: val.Source.AuthorEmail,
			CreateTime:  timestamppb.New(val.Source.CreateTime),
			QuickRead:   val.Source.QuickRead,
		}
		resp = append(resp, res)
	}

	return resp
}

func partialOrAllUpdate(method string, existingArt *pb.GetArticleByIdResp, reqArt *pb.EditArticleReq) *pb.EditArticleReq {
	procdArt := &pb.EditArticleReq{Id: reqArt.Id}
	log.Printf("existingArt tags: %+v", existingArt)
	log.Println("Requested tags: ", reqArt.Tags)

	if method == http.MethodPatch {
		if reqArt.Title == "" {
			procdArt.Title = existingArt.Title
		} else {
			procdArt.Title = reqArt.Title
		}
		if reqArt.Content == "" {
			procdArt.Content = existingArt.Content
		} else {
			procdArt.Content = reqArt.Content
		}
		if len(reqArt.Tags) == 0 {
			procdArt.Tags = existingArt.Tags
		} else {
			procdArt.Tags = reqArt.Tags
		}
	} else {
		procdArt.Title = reqArt.Title
		procdArt.Content = reqArt.Content
		procdArt.Tags = reqArt.Tags
	}
	log.Println("TobeUPdated tags: ", procdArt.Tags)

	return procdArt
}
