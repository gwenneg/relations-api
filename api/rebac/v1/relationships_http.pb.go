// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: rebac/v1/relationships.proto

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

const OperationRelationshipsCreateRelationships = "/api.rebac.v1.Relationships/CreateRelationships"
const OperationRelationshipsDeleteRelationships = "/api.rebac.v1.Relationships/DeleteRelationships"
const OperationRelationshipsReadRelationships = "/api.rebac.v1.Relationships/ReadRelationships"

type RelationshipsHTTPServer interface {
	CreateRelationships(context.Context, *CreateRelationshipsRequest) (*CreateRelationshipsResponse, error)
	DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error)
	ReadRelationships(context.Context, *ReadRelationshipsRequest) (*ReadRelationshipsResponse, error)
}

func RegisterRelationshipsHTTPServer(s *http.Server, srv RelationshipsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/relationships", _Relationships_CreateRelationships0_HTTP_Handler(srv))
	r.GET("/v1/relationships", _Relationships_ReadRelationships0_HTTP_Handler(srv))
	r.DELETE("/v1/relationships", _Relationships_DeleteRelationships0_HTTP_Handler(srv))
}

func _Relationships_CreateRelationships0_HTTP_Handler(srv RelationshipsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRelationshipsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationshipsCreateRelationships)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRelationships(ctx, req.(*CreateRelationshipsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRelationshipsResponse)
		return ctx.Result(200, reply)
	}
}

func _Relationships_ReadRelationships0_HTTP_Handler(srv RelationshipsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReadRelationshipsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationshipsReadRelationships)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReadRelationships(ctx, req.(*ReadRelationshipsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReadRelationshipsResponse)
		return ctx.Result(200, reply)
	}
}

func _Relationships_DeleteRelationships0_HTTP_Handler(srv RelationshipsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRelationshipsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationshipsDeleteRelationships)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRelationships(ctx, req.(*DeleteRelationshipsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteRelationshipsResponse)
		return ctx.Result(200, reply)
	}
}

type RelationshipsHTTPClient interface {
	CreateRelationships(ctx context.Context, req *CreateRelationshipsRequest, opts ...http.CallOption) (rsp *CreateRelationshipsResponse, err error)
	DeleteRelationships(ctx context.Context, req *DeleteRelationshipsRequest, opts ...http.CallOption) (rsp *DeleteRelationshipsResponse, err error)
	ReadRelationships(ctx context.Context, req *ReadRelationshipsRequest, opts ...http.CallOption) (rsp *ReadRelationshipsResponse, err error)
}

type RelationshipsHTTPClientImpl struct {
	cc *http.Client
}

func NewRelationshipsHTTPClient(client *http.Client) RelationshipsHTTPClient {
	return &RelationshipsHTTPClientImpl{client}
}

func (c *RelationshipsHTTPClientImpl) CreateRelationships(ctx context.Context, in *CreateRelationshipsRequest, opts ...http.CallOption) (*CreateRelationshipsResponse, error) {
	var out CreateRelationshipsResponse
	pattern := "/v1/relationships"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRelationshipsCreateRelationships))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipsHTTPClientImpl) DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...http.CallOption) (*DeleteRelationshipsResponse, error) {
	var out DeleteRelationshipsResponse
	pattern := "/v1/relationships"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRelationshipsDeleteRelationships))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipsHTTPClientImpl) ReadRelationships(ctx context.Context, in *ReadRelationshipsRequest, opts ...http.CallOption) (*ReadRelationshipsResponse, error) {
	var out ReadRelationshipsResponse
	pattern := "/v1/relationships"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRelationshipsReadRelationships))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
