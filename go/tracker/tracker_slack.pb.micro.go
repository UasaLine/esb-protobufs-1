// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tracker_slack.proto

package tracker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for TrackerSlack service

func NewTrackerSlackEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TrackerSlack service

type TrackerSlackService interface {
	SendTicket(ctx context.Context, in *SendTicketParams, opts ...client.CallOption) (*SendTicketResponse, error)
}

type trackerSlackService struct {
	c    client.Client
	name string
}

func NewTrackerSlackService(name string, c client.Client) TrackerSlackService {
	return &trackerSlackService{
		c:    c,
		name: name,
	}
}

func (c *trackerSlackService) SendTicket(ctx context.Context, in *SendTicketParams, opts ...client.CallOption) (*SendTicketResponse, error) {
	req := c.c.NewRequest(c.name, "TrackerSlack.SendTicket", in)
	out := new(SendTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrackerSlack service

type TrackerSlackHandler interface {
	SendTicket(context.Context, *SendTicketParams, *SendTicketResponse) error
}

func RegisterTrackerSlackHandler(s server.Server, hdlr TrackerSlackHandler, opts ...server.HandlerOption) error {
	type trackerSlack interface {
		SendTicket(ctx context.Context, in *SendTicketParams, out *SendTicketResponse) error
	}
	type TrackerSlack struct {
		trackerSlack
	}
	h := &trackerSlackHandler{hdlr}
	return s.Handle(s.NewHandler(&TrackerSlack{h}, opts...))
}

type trackerSlackHandler struct {
	TrackerSlackHandler
}

func (h *trackerSlackHandler) SendTicket(ctx context.Context, in *SendTicketParams, out *SendTicketResponse) error {
	return h.TrackerSlackHandler.SendTicket(ctx, in, out)
}
