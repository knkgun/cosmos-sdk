// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vestingv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error) {
	out := new(MsgCreateVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error) {
	out := new(MsgCreatePeriodicVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVestingAccount not implemented")
}
func (*UnimplementedMsgServer) CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeriodicVestingAccount not implemented")
}
func (*UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

func RegisterMsgServer(s *grpc.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVestingAccount(ctx, req.(*MsgCreateVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePeriodicVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeriodicVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, req.(*MsgCreatePeriodicVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vesting.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVestingAccount",
			Handler:    _Msg_CreateVestingAccount_Handler,
		},
		{
			MethodName: "CreatePeriodicVestingAccount",
			Handler:    _Msg_CreatePeriodicVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/vesting/v1beta1/tx.proto",
}
