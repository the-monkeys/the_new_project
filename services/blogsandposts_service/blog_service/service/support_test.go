package service

import (
	"reflect"
	"testing"

	"github.com/89minutes/the_new_project/services/blogsandposts_service/blog_service/pb"
)

func Test_partialOrAllUpdate(t *testing.T) {
	type args struct {
		isPartial   bool
		existingArt *pb.GetBlogByIdResponse
		reqArt      *pb.EditBlogRequest
	}
	tests := []struct {
		name string
		args args
		want *pb.EditBlogRequest
	}{
		// TODO: Update the params
		{
			name: "Update only title",
			args: args{
				isPartial:   true,
				existingArt: &pb.GetBlogByIdResponse{},
				reqArt:      &pb.EditBlogRequest{},
			},
			want: &pb.EditBlogRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partialOrAllUpdate(tt.args.isPartial, tt.args.existingArt, tt.args.reqArt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partialOrAllUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
