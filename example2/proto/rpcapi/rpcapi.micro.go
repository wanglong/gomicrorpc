// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: example2/proto/rpcapi/rpcapi.proto

/*
Package rpcapi is a generated protocol buffer package.

It is generated from these files:
	example2/proto/rpcapi/rpcapi.proto

It has these top-level messages:
*/
package rpcapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/lpxxn/gomicrorpc/example2/proto/model"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = model.SResponse{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayService interface {
	Hello(ctx context.Context, in *model.SayParam, opts ...client.CallOption) (*model.SayResponse, error)
	Stream(ctx context.Context, in *model.SRequest, opts ...client.CallOption) (Say_StreamService, error)
}

type sayService struct {
	c    client.Client
	name string
}

func NewSayService(name string, c client.Client) SayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "rpcapi"
	}
	return &sayService{
		c:    c,
		name: name,
	}
}

func (c *sayService) Hello(ctx context.Context, in *model.SayParam, opts ...client.CallOption) (*model.SayResponse, error) {
	req := c.c.NewRequest(c.name, "Say.Hello", in)
	out := new(model.SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sayService) Stream(ctx context.Context, in *model.SRequest, opts ...client.CallOption) (Say_StreamService, error) {
	req := c.c.NewRequest(c.name, "Say.Stream", &model.SRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &sayServiceStream{stream}, nil
}

type Say_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*model.SResponse, error)
}

type sayServiceStream struct {
	stream client.Stream
}

func (x *sayServiceStream) Close() error {
	return x.stream.Close()
}

func (x *sayServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sayServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sayServiceStream) Recv() (*model.SResponse, error) {
	m := new(model.SResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *model.SayParam, *model.SayResponse) error
	Stream(context.Context, *model.SRequest, Say_StreamStream) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) {
	type say interface {
		Hello(ctx context.Context, in *model.SayParam, out *model.SayResponse) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Say struct {
		say
	}
	h := &sayHandler{hdlr}
	s.Handle(s.NewHandler(&Say{h}, opts...))
}

type sayHandler struct {
	SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, in *model.SayParam, out *model.SayResponse) error {
	return h.SayHandler.Hello(ctx, in, out)
}

func (h *sayHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(model.SRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SayHandler.Stream(ctx, m, &sayStreamStream{stream})
}

type Say_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*model.SResponse) error
}

type sayStreamStream struct {
	stream server.Stream
}

func (x *sayStreamStream) Close() error {
	return x.stream.Close()
}

func (x *sayStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sayStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sayStreamStream) Send(m *model.SResponse) error {
	return x.stream.Send(m)
}
