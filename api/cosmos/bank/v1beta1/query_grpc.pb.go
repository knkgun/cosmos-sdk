// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bankv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Balance queries the balance of a single coin for a single account.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// AllBalances queries the balance of all coins for a single account.
	AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error)
	// TotalSupply queries the total supply of all coins.
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)
	// SupplyOf queries the supply of a single coin.
	SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error)
	// Params queries the parameters of x/bank module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// DenomsMetadata queries the client metadata of a given coin denomination.
	DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error)
	// DenomsMetadata queries the client metadata for all registered coin
	// denominations.
	DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error)
	// DenomOwners queries for all account addresses that own a particular token
	// denomination.
	DenomOwners(ctx context.Context, in *QueryDenomOwnersRequest, opts ...grpc.CallOption) (*QueryDenomOwnersResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error) {
	out := new(QueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/AllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error) {
	out := new(QuerySupplyOfResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/SupplyOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error) {
	out := new(QueryDenomMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/DenomMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error) {
	out := new(QueryDenomsMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/DenomsMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomOwners(ctx context.Context, in *QueryDenomOwnersRequest, opts ...grpc.CallOption) (*QueryDenomOwnersResponse, error) {
	out := new(QueryDenomOwnersResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/DenomOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Balance queries the balance of a single coin for a single account.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// AllBalances queries the balance of all coins for a single account.
	AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)
	// TotalSupply queries the total supply of all coins.
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)
	// SupplyOf queries the supply of a single coin.
	SupplyOf(context.Context, *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error)
	// Params queries the parameters of x/bank module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// DenomsMetadata queries the client metadata of a given coin denomination.
	DenomMetadata(context.Context, *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error)
	// DenomsMetadata queries the client metadata for all registered coin
	// denominations.
	DenomsMetadata(context.Context, *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error)
	// DenomOwners queries for all account addresses that own a particular token
	// denomination.
	DenomOwners(context.Context, *QueryDenomOwnersRequest) (*QueryDenomOwnersResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (*UnimplementedQueryServer) AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllBalances not implemented")
}
func (*UnimplementedQueryServer) TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) SupplyOf(context.Context, *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplyOf not implemented")
}
func (*UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomMetadata(context.Context, *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomMetadata not implemented")
}
func (*UnimplementedQueryServer) DenomsMetadata(context.Context, *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomsMetadata not implemented")
}
func (*UnimplementedQueryServer) DenomOwners(context.Context, *QueryDenomOwnersRequest) (*QueryDenomOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomOwners not implemented")
}
func (*UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/AllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllBalances(ctx, req.(*QueryAllBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SupplyOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SupplyOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/SupplyOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SupplyOf(ctx, req.(*QuerySupplyOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/DenomMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomMetadata(ctx, req.(*QueryDenomMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/DenomsMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomsMetadata(ctx, req.(*QueryDenomsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/DenomOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomOwners(ctx, req.(*QueryDenomOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.bank.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "AllBalances",
			Handler:    _Query_AllBalances_Handler,
		},
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "SupplyOf",
			Handler:    _Query_SupplyOf_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomMetadata",
			Handler:    _Query_DenomMetadata_Handler,
		},
		{
			MethodName: "DenomsMetadata",
			Handler:    _Query_DenomsMetadata_Handler,
		},
		{
			MethodName: "DenomOwners",
			Handler:    _Query_DenomOwners_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/bank/v1beta1/query.proto",
}
