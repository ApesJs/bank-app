// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: service_bank_app.proto

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

// BankAppClient is the client API for BankApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankAppClient interface {
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResult, error)
}

type bankAppClient struct {
	cc grpc.ClientConnInterface
}

func NewBankAppClient(cc grpc.ClientConnInterface) BankAppClient {
	return &bankAppClient{cc}
}

func (c *bankAppClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResult, error) {
	out := new(CreateTransferResult)
	err := c.cc.Invoke(ctx, "/pb.BankApp/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankAppServer is the server API for BankApp service.
// All implementations must embed UnimplementedBankAppServer
// for forward compatibility
type BankAppServer interface {
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResult, error)
	mustEmbedUnimplementedBankAppServer()
}

// UnimplementedBankAppServer must be embedded to have forward compatible implementations.
type UnimplementedBankAppServer struct {
}

func (UnimplementedBankAppServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedBankAppServer) mustEmbedUnimplementedBankAppServer() {}

// UnsafeBankAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankAppServer will
// result in compilation errors.
type UnsafeBankAppServer interface {
	mustEmbedUnimplementedBankAppServer()
}

func RegisterBankAppServer(s grpc.ServiceRegistrar, srv BankAppServer) {
	s.RegisterService(&BankApp_ServiceDesc, srv)
}

func _BankApp_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAppServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BankApp/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAppServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankApp_ServiceDesc is the grpc.ServiceDesc for BankApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BankApp",
	HandlerType: (*BankAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransfer",
			Handler:    _BankApp_CreateTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_bank_app.proto",
}
