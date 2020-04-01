// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/slack.proto

package slack

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Tracker service

type TrackerService interface {
	StatusChangeTicket(ctx context.Context, in *StatusChangeTicketParams, opts ...client.CallOption) (*StatusChangeTicketResponse, error)
}

type trackerService struct {
	c    client.Client
	name string
}

func NewTrackerService(name string, c client.Client) TrackerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "slack"
	}
	return &trackerService{
		c:    c,
		name: name,
	}
}

func (c *trackerService) StatusChangeTicket(ctx context.Context, in *StatusChangeTicketParams, opts ...client.CallOption) (*StatusChangeTicketResponse, error) {
	req := c.c.NewRequest(c.name, "Tracker.StatusChangeTicket", in)
	out := new(StatusChangeTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tracker service

type TrackerHandler interface {
	StatusChangeTicket(context.Context, *StatusChangeTicketParams, *StatusChangeTicketResponse) error
}

func RegisterTrackerHandler(s server.Server, hdlr TrackerHandler, opts ...server.HandlerOption) error {
	type tracker interface {
		StatusChangeTicket(ctx context.Context, in *StatusChangeTicketParams, out *StatusChangeTicketResponse) error
	}
	type Tracker struct {
		tracker
	}
	h := &trackerHandler{hdlr}
	return s.Handle(s.NewHandler(&Tracker{h}, opts...))
}

type trackerHandler struct {
	TrackerHandler
}

func (h *trackerHandler) StatusChangeTicket(ctx context.Context, in *StatusChangeTicketParams, out *StatusChangeTicketResponse) error {
	return h.TrackerHandler.StatusChangeTicket(ctx, in, out)
}