// Code generated by protoc-gen-go.
// source: twitter.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	twitter.proto

It has these top-level messages:
	Ack
	User
	Tweet
	Search
	Timeline
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type Ack struct {
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto1.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}

type User struct {
	ID uint64 `protobuf:"varint,1,opt" json:"ID,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()    {}

type Tweet struct {
	ID   uint64 `protobuf:"varint,1,opt" json:"ID,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	User *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *Tweet) Reset()         { *m = Tweet{} }
func (m *Tweet) String() string { return proto1.CompactTextString(m) }
func (*Tweet) ProtoMessage()    {}

func (m *Tweet) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Search struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Search) Reset()         { *m = Search{} }
func (m *Search) String() string { return proto1.CompactTextString(m) }
func (*Search) ProtoMessage()    {}

type Timeline struct {
	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets" json:"tweets,omitempty"`
}

func (m *Timeline) Reset()         { *m = Timeline{} }
func (m *Timeline) String() string { return proto1.CompactTextString(m) }
func (*Timeline) ProtoMessage()    {}

func (m *Timeline) GetTweets() []*Tweet {
	if m != nil {
		return m.Tweets
	}
	return nil
}

func init() {
}

// Client API for Twitter service

type TwitterClient interface {
	GetTimeline(ctx context.Context, in *User, opts ...grpc.CallOption) (*Timeline, error)
	Firehose(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_FirehoseClient, error)
	Add(ctx context.Context, in *Tweet, opts ...grpc.CallOption) (*Ack, error)
}

type twitterClient struct {
	cc *grpc.ClientConn
}

func NewTwitterClient(cc *grpc.ClientConn) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) GetTimeline(ctx context.Context, in *User, opts ...grpc.CallOption) (*Timeline, error) {
	out := new(Timeline)
	err := grpc.Invoke(ctx, "/proto.Twitter/GetTimeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) Firehose(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_FirehoseClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Twitter_serviceDesc.Streams[0], c.cc, "/proto.Twitter/Firehose", opts...)
	if err != nil {
		return nil, err
	}
	x := &twitterFirehoseClient{stream}
	if err := x.ClientStream.SendProto(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Twitter_FirehoseClient interface {
	Recv() (*Tweet, error)
	grpc.ClientStream
}

type twitterFirehoseClient struct {
	grpc.ClientStream
}

func (x *twitterFirehoseClient) Recv() (*Tweet, error) {
	m := new(Tweet)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *twitterClient) Add(ctx context.Context, in *Tweet, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/proto.Twitter/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Twitter service

type TwitterServer interface {
	GetTimeline(context.Context, *User) (*Timeline, error)
	Firehose(*Search, Twitter_FirehoseServer) error
	Add(context.Context, *Tweet) (*Ack, error)
}

func RegisterTwitterServer(s *grpc.Server, srv TwitterServer) {
	s.RegisterService(&_Twitter_serviceDesc, srv)
}

func _Twitter_GetTimeline_Handler(srv interface{}, ctx context.Context, buf []byte) (proto1.Message, error) {
	in := new(User)
	if err := proto1.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TwitterServer).GetTimeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Twitter_Firehose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Search)
	if err := stream.RecvProto(m); err != nil {
		return err
	}
	return srv.(TwitterServer).Firehose(m, &twitterFirehoseServer{stream})
}

type Twitter_FirehoseServer interface {
	Send(*Tweet) error
	grpc.ServerStream
}

type twitterFirehoseServer struct {
	grpc.ServerStream
}

func (x *twitterFirehoseServer) Send(m *Tweet) error {
	return x.ServerStream.SendProto(m)
}

func _Twitter_Add_Handler(srv interface{}, ctx context.Context, buf []byte) (proto1.Message, error) {
	in := new(Tweet)
	if err := proto1.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TwitterServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Twitter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Twitter",
	HandlerType: (*TwitterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeline",
			Handler:    _Twitter_GetTimeline_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Twitter_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Firehose",
			Handler:       _Twitter_Firehose_Handler,
			ServerStreams: true,
		},
	},
}
