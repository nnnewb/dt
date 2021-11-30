// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error)
	GetWallet(ctx context.Context, in *GetWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error) {
	out := new(ListWalletsResponse)
	err := c.cc.Invoke(ctx, "/WalletService/ListWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetWallet(ctx context.Context, in *GetWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/WalletService/GetWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/WalletService/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/WalletService/UpdateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WalletService/DeleteWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error)
	GetWallet(context.Context, *GetWalletRequest) (*Wallet, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error)
	UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error)
	DeleteWallet(context.Context, *DeleteWalletRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallets not implemented")
}
func (UnimplementedWalletServiceServer) GetWallet(context.Context, *GetWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallet not implemented")
}
func (UnimplementedWalletServiceServer) CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServiceServer) UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWallet not implemented")
}
func (UnimplementedWalletServiceServer) DeleteWallet(context.Context, *DeleteWalletRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWallet not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_ListWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).ListWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WalletService/ListWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).ListWallets(ctx, req.(*ListWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WalletService/GetWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWallet(ctx, req.(*GetWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WalletService/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_UpdateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UpdateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WalletService/UpdateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UpdateWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_DeleteWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).DeleteWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WalletService/DeleteWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).DeleteWallet(ctx, req.(*DeleteWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWallets",
			Handler:    _WalletService_ListWallets_Handler,
		},
		{
			MethodName: "GetWallet",
			Handler:    _WalletService_GetWallet_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _WalletService_CreateWallet_Handler,
		},
		{
			MethodName: "UpdateWallet",
			Handler:    _WalletService_UpdateWallet_Handler,
		},
		{
			MethodName: "DeleteWallet",
			Handler:    _WalletService_DeleteWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}