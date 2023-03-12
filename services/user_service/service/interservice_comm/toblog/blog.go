package toblog

import (
	"context"
	"log"

	"github.com/89minutes/the_new_project/apis/interservice/blogs/pb"
	"google.golang.org/grpc"
)

type BlogClient struct {
	blogServiceClient pb.BlogServiceClient
}

func NewClient(conn *grpc.ClientConn) *BlogClient {
	return &BlogClient{blogServiceClient: pb.NewBlogServiceClient(conn)}
}

// TODO: Update error handling
func (bc *BlogClient) UpdateBlogsUserDeactivated(conn *grpc.ClientConn) {
	// send a request to the server
	res, err := bc.blogServiceClient.SetUserDeactivated(context.Background(), &pb.SetUserDeactivatedReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// print the response from the server
	log.Printf("Greeting: %s", res.Message)
}
