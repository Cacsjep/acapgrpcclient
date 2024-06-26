// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.2
// source: cctv.proto

package cctv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CCTVService_StreamVideo_FullMethodName = "/CCTVService/StreamVideo"
)

// CCTVServiceClient is the client API for CCTVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CCTVServiceClient interface {
	StreamVideo(ctx context.Context, opts ...grpc.CallOption) (CCTVService_StreamVideoClient, error)
}

type cCTVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCCTVServiceClient(cc grpc.ClientConnInterface) CCTVServiceClient {
	return &cCTVServiceClient{cc}
}

func (c *cCTVServiceClient) StreamVideo(ctx context.Context, opts ...grpc.CallOption) (CCTVService_StreamVideoClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CCTVService_ServiceDesc.Streams[0], CCTVService_StreamVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &cCTVServiceStreamVideoClient{ClientStream: stream}
	return x, nil
}

type CCTVService_StreamVideoClient interface {
	Send(*FrameMessage) error
	Recv() (*Acknowledgement, error)
	grpc.ClientStream
}

type cCTVServiceStreamVideoClient struct {
	grpc.ClientStream
}

func (x *cCTVServiceStreamVideoClient) Send(m *FrameMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cCTVServiceStreamVideoClient) Recv() (*Acknowledgement, error) {
	m := new(Acknowledgement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CCTVServiceServer is the server API for CCTVService service.
// All implementations must embed UnimplementedCCTVServiceServer
// for forward compatibility
type CCTVServiceServer interface {
	StreamVideo(CCTVService_StreamVideoServer) error
	mustEmbedUnimplementedCCTVServiceServer()
}

// UnimplementedCCTVServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCCTVServiceServer struct {
}

func (UnimplementedCCTVServiceServer) StreamVideo(CCTVService_StreamVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}
func (UnimplementedCCTVServiceServer) mustEmbedUnimplementedCCTVServiceServer() {}

// UnsafeCCTVServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CCTVServiceServer will
// result in compilation errors.
type UnsafeCCTVServiceServer interface {
	mustEmbedUnimplementedCCTVServiceServer()
}

func RegisterCCTVServiceServer(s grpc.ServiceRegistrar, srv CCTVServiceServer) {
	s.RegisterService(&CCTVService_ServiceDesc, srv)
}

func _CCTVService_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CCTVServiceServer).StreamVideo(&cCTVServiceStreamVideoServer{ServerStream: stream})
}

type CCTVService_StreamVideoServer interface {
	Send(*Acknowledgement) error
	Recv() (*FrameMessage, error)
	grpc.ServerStream
}

type cCTVServiceStreamVideoServer struct {
	grpc.ServerStream
}

func (x *cCTVServiceStreamVideoServer) Send(m *Acknowledgement) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cCTVServiceStreamVideoServer) Recv() (*FrameMessage, error) {
	m := new(FrameMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CCTVService_ServiceDesc is the grpc.ServiceDesc for CCTVService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CCTVService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CCTVService",
	HandlerType: (*CCTVServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamVideo",
			Handler:       _CCTVService_StreamVideo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cctv.proto",
}
