// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: services/api_gateway/pkg/blogsandposts/pb/blogsandposts.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlogsAndPostServiceClient is the client API for BlogsAndPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogsAndPostServiceClient interface {
	CreateABlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
}

type blogsAndPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogsAndPostServiceClient(cc grpc.ClientConnInterface) BlogsAndPostServiceClient {
	return &blogsAndPostServiceClient{cc}
}

func (c *blogsAndPostServiceClient) CreateABlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := c.cc.Invoke(ctx, "/auth.BlogsAndPostService/CreateABlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogsAndPostServiceServer is the server API for BlogsAndPostService service.
// All implementations must embed UnimplementedBlogsAndPostServiceServer
// for forward compatibility
type BlogsAndPostServiceServer interface {
	CreateABlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	mustEmbedUnimplementedBlogsAndPostServiceServer()
}

// UnimplementedBlogsAndPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogsAndPostServiceServer struct {
}

func (UnimplementedBlogsAndPostServiceServer) CreateABlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateABlog not implemented")
}
func (UnimplementedBlogsAndPostServiceServer) mustEmbedUnimplementedBlogsAndPostServiceServer() {}

// UnsafeBlogsAndPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogsAndPostServiceServer will
// result in compilation errors.
type UnsafeBlogsAndPostServiceServer interface {
	mustEmbedUnimplementedBlogsAndPostServiceServer()
}

func RegisterBlogsAndPostServiceServer(s grpc.ServiceRegistrar, srv BlogsAndPostServiceServer) {
	s.RegisterService(&BlogsAndPostService_ServiceDesc, srv)
}

func _BlogsAndPostService_CreateABlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsAndPostServiceServer).CreateABlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.BlogsAndPostService/CreateABlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsAndPostServiceServer).CreateABlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogsAndPostService_ServiceDesc is the grpc.ServiceDesc for BlogsAndPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogsAndPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.BlogsAndPostService",
	HandlerType: (*BlogsAndPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateABlog",
			Handler:    _BlogsAndPostService_CreateABlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/api_gateway/pkg/blogsandposts/pb/blogsandposts.proto",
}
