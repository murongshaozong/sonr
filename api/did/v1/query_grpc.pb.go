// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: did/v1/query.proto

package didv1

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

const (
	Query_Params_FullMethodName                   = "/did.v1.Query/Params"
	Query_ParamsAssets_FullMethodName             = "/did.v1.Query/ParamsAssets"
	Query_ParamsByAsset_FullMethodName            = "/did.v1.Query/ParamsByAsset"
	Query_ParamsKeys_FullMethodName               = "/did.v1.Query/ParamsKeys"
	Query_ParamsByKey_FullMethodName              = "/did.v1.Query/ParamsByKey"
	Query_RegistrationOptionsByKey_FullMethodName = "/did.v1.Query/RegistrationOptionsByKey"
	Query_Resolve_FullMethodName                  = "/did.v1.Query/Resolve"
	Query_Service_FullMethodName                  = "/did.v1.Query/Service"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries all parameters of the module.
	Params(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ParamsAssets queries all parameters of the module.
	ParamsAssets(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Params queries all parameters of the module.
	ParamsByAsset(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// ParamsKeys queries all parameters of the module.
	ParamsKeys(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Params queries all parameters of the module.
	ParamsByKey(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Params queries all parameters of the module.
	RegistrationOptionsByKey(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Resolve queries the DID document by its id.
	Resolve(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Service returns associated ServiceInfo for a given Origin
	// if the servie is not found, a fingerprint is generated to be used
	// as a TXT record in DNS. v=sonr, o=origin, p=protocol
	Service(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ParamsAssets(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_ParamsAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ParamsByAsset(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_ParamsByAsset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ParamsKeys(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_ParamsKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ParamsByKey(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_ParamsByKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RegistrationOptionsByKey(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_RegistrationOptionsByKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Resolve(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_Resolve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Service(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Query_Service_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries all parameters of the module.
	Params(context.Context, *QueryRequest) (*QueryParamsResponse, error)
	// ParamsAssets queries all parameters of the module.
	ParamsAssets(context.Context, *QueryRequest) (*QueryResponse, error)
	// Params queries all parameters of the module.
	ParamsByAsset(context.Context, *QueryRequest) (*QueryResponse, error)
	// ParamsKeys queries all parameters of the module.
	ParamsKeys(context.Context, *QueryRequest) (*QueryResponse, error)
	// Params queries all parameters of the module.
	ParamsByKey(context.Context, *QueryRequest) (*QueryResponse, error)
	// Params queries all parameters of the module.
	RegistrationOptionsByKey(context.Context, *QueryRequest) (*QueryResponse, error)
	// Resolve queries the DID document by its id.
	Resolve(context.Context, *QueryRequest) (*QueryResponse, error)
	// Service returns associated ServiceInfo for a given Origin
	// if the servie is not found, a fingerprint is generated to be used
	// as a TXT record in DNS. v=sonr, o=origin, p=protocol
	Service(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ParamsAssets(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParamsAssets not implemented")
}
func (UnimplementedQueryServer) ParamsByAsset(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParamsByAsset not implemented")
}
func (UnimplementedQueryServer) ParamsKeys(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParamsKeys not implemented")
}
func (UnimplementedQueryServer) ParamsByKey(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParamsByKey not implemented")
}
func (UnimplementedQueryServer) RegistrationOptionsByKey(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationOptionsByKey not implemented")
}
func (UnimplementedQueryServer) Resolve(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedQueryServer) Service(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Service not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ParamsAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParamsAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ParamsAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParamsAssets(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ParamsByAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParamsByAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ParamsByAsset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParamsByAsset(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ParamsKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParamsKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ParamsKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParamsKeys(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ParamsByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParamsByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ParamsByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParamsByKey(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RegistrationOptionsByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegistrationOptionsByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RegistrationOptionsByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegistrationOptionsByKey(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Resolve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Resolve(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Service_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Service(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Service_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Service(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "did.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ParamsAssets",
			Handler:    _Query_ParamsAssets_Handler,
		},
		{
			MethodName: "ParamsByAsset",
			Handler:    _Query_ParamsByAsset_Handler,
		},
		{
			MethodName: "ParamsKeys",
			Handler:    _Query_ParamsKeys_Handler,
		},
		{
			MethodName: "ParamsByKey",
			Handler:    _Query_ParamsByKey_Handler,
		},
		{
			MethodName: "RegistrationOptionsByKey",
			Handler:    _Query_RegistrationOptionsByKey_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _Query_Resolve_Handler,
		},
		{
			MethodName: "Service",
			Handler:    _Query_Service_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/v1/query.proto",
}
