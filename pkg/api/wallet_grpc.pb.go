// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc1
// source: wallet.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Wallets_CreateWallet_FullMethodName  = "/wallet_proto.Wallets/CreateWallet"
	Wallets_GetBalance_FullMethodName    = "/wallet_proto.Wallets/GetBalance"
	Wallets_UpdateBalance_FullMethodName = "/wallet_proto.Wallets/UpdateBalance"
)

// WalletsClient is the client API for Wallets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletsClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	UpdateBalance(ctx context.Context, in *UpdateBalaceRequest, opts ...grpc.CallOption) (*UpdateBalaceResponse, error)
}

type walletsClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletsClient(cc grpc.ClientConnInterface) WalletsClient {
	return &walletsClient{cc}
}

func (c *walletsClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, Wallets_CreateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, Wallets_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsClient) UpdateBalance(ctx context.Context, in *UpdateBalaceRequest, opts ...grpc.CallOption) (*UpdateBalaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBalaceResponse)
	err := c.cc.Invoke(ctx, Wallets_UpdateBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletsServer is the server API for Wallets service.
// All implementations must embed UnimplementedWalletsServer
// for forward compatibility.
type WalletsServer interface {
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	UpdateBalance(context.Context, *UpdateBalaceRequest) (*UpdateBalaceResponse, error)
	mustEmbedUnimplementedWalletsServer()
}

// UnimplementedWalletsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletsServer struct{}

func (UnimplementedWalletsServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletsServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedWalletsServer) UpdateBalance(context.Context, *UpdateBalaceRequest) (*UpdateBalaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedWalletsServer) mustEmbedUnimplementedWalletsServer() {}
func (UnimplementedWalletsServer) testEmbeddedByValue()                 {}

// UnsafeWalletsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletsServer will
// result in compilation errors.
type UnsafeWalletsServer interface {
	mustEmbedUnimplementedWalletsServer()
}

func RegisterWalletsServer(s grpc.ServiceRegistrar, srv WalletsServer) {
	// If the following call pancis, it indicates UnimplementedWalletsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Wallets_ServiceDesc, srv)
}

func _Wallets_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallets_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallets_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallets_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallets_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallets_UpdateBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).UpdateBalance(ctx, req.(*UpdateBalaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallets_ServiceDesc is the grpc.ServiceDesc for Wallets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet_proto.Wallets",
	HandlerType: (*WalletsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _Wallets_CreateWallet_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Wallets_GetBalance_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _Wallets_UpdateBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
