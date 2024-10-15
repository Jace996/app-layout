// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             (unknown)
// source: api/post/v1/post.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPostServiceCreatePost = "/kit_layout.api.post.v1.PostService/CreatePost"
const OperationPostServiceDeletePost = "/kit_layout.api.post.v1.PostService/DeletePost"
const OperationPostServiceGetPost = "/kit_layout.api.post.v1.PostService/GetPost"
const OperationPostServiceListPost = "/kit_layout.api.post.v1.PostService/ListPost"
const OperationPostServiceUpdatePost = "/kit_layout.api.post.v1.PostService/UpdatePost"

type PostServiceHTTPServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	ListPost(context.Context, *ListPostRequest) (*ListPostReply, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
}

func RegisterPostServiceHTTPServer(s *http.Server, srv PostServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/github.com/jace996/app-layout/post/list", _PostService_ListPost0_HTTP_Handler(srv))
	r.GET("/v1/github.com/jace996/app-layout/posts", _PostService_ListPost1_HTTP_Handler(srv))
	r.GET("/v1/github.com/jace996/app-layout/post/{id}", _PostService_GetPost0_HTTP_Handler(srv))
	r.POST("/v1/github.com/jace996/app-layout/post", _PostService_CreatePost0_HTTP_Handler(srv))
	r.PATCH("/v1/github.com/jace996/app-layout/post/{post.id}", _PostService_UpdatePost0_HTTP_Handler(srv))
	r.PUT("/v1/github.com/jace996/app-layout/post/{post.id}", _PostService_UpdatePost1_HTTP_Handler(srv))
	r.DELETE("/v1/github.com/jace996/app-layout/post/{id}", _PostService_DeletePost0_HTTP_Handler(srv))
}

func _PostService_ListPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceListPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPost(ctx, req.(*ListPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPostReply)
		return ctx.Result(200, reply)
	}
}

func _PostService_ListPost1_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceListPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPost(ctx, req.(*ListPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPostReply)
		return ctx.Result(200, reply)
	}
}

func _PostService_GetPost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceGetPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPost(ctx, req.(*GetPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_CreatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceCreatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePost(ctx, req.(*CreatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_UpdatePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_UpdatePost1_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _PostService_DeletePost0_HTTP_Handler(srv PostServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPostServiceDeletePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePost(ctx, req.(*DeletePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePostReply)
		return ctx.Result(200, reply)
	}
}

type PostServiceHTTPClient interface {
	CreatePost(ctx context.Context, req *CreatePostRequest, opts ...http.CallOption) (rsp *Post, err error)
	DeletePost(ctx context.Context, req *DeletePostRequest, opts ...http.CallOption) (rsp *DeletePostReply, err error)
	GetPost(ctx context.Context, req *GetPostRequest, opts ...http.CallOption) (rsp *Post, err error)
	ListPost(ctx context.Context, req *ListPostRequest, opts ...http.CallOption) (rsp *ListPostReply, err error)
	UpdatePost(ctx context.Context, req *UpdatePostRequest, opts ...http.CallOption) (rsp *Post, err error)
}

type PostServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPostServiceHTTPClient(client *http.Client) PostServiceHTTPClient {
	return &PostServiceHTTPClientImpl{client}
}

func (c *PostServiceHTTPClientImpl) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/v1/github.com/jace996/app-layout/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceCreatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostServiceHTTPClientImpl) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...http.CallOption) (*DeletePostReply, error) {
	var out DeletePostReply
	pattern := "/v1/github.com/jace996/app-layout/post/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceDeletePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostServiceHTTPClientImpl) GetPost(ctx context.Context, in *GetPostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/v1/github.com/jace996/app-layout/post/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceGetPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostServiceHTTPClientImpl) ListPost(ctx context.Context, in *ListPostRequest, opts ...http.CallOption) (*ListPostReply, error) {
	var out ListPostReply
	pattern := "/v1/github.com/jace996/app-layout/posts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPostServiceListPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PostServiceHTTPClientImpl) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/v1/github.com/jace996/app-layout/post/{post.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPostServiceUpdatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
