// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	// TabularDataByFilter queries tabular data and metadata based on given filters.
	TabularDataByFilter(ctx context.Context, in *TabularDataByFilterRequest, opts ...grpc.CallOption) (*TabularDataByFilterResponse, error)
	// BinaryDataByFilter queries binary data and metadata based on given filters.
	BinaryDataByFilter(ctx context.Context, in *BinaryDataByFilterRequest, opts ...grpc.CallOption) (*BinaryDataByFilterResponse, error)
	// BinaryDataByIDs queries binary data and metadata based on given IDs.
	BinaryDataByIDs(ctx context.Context, in *BinaryDataByIDsRequest, opts ...grpc.CallOption) (*BinaryDataByIDsResponse, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) TabularDataByFilter(ctx context.Context, in *TabularDataByFilterRequest, opts ...grpc.CallOption) (*TabularDataByFilterResponse, error) {
	out := new(TabularDataByFilterResponse)
	err := c.cc.Invoke(ctx, "/viam.app.data.v1.DataService/TabularDataByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) BinaryDataByFilter(ctx context.Context, in *BinaryDataByFilterRequest, opts ...grpc.CallOption) (*BinaryDataByFilterResponse, error) {
	out := new(BinaryDataByFilterResponse)
	err := c.cc.Invoke(ctx, "/viam.app.data.v1.DataService/BinaryDataByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) BinaryDataByIDs(ctx context.Context, in *BinaryDataByIDsRequest, opts ...grpc.CallOption) (*BinaryDataByIDsResponse, error) {
	out := new(BinaryDataByIDsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.data.v1.DataService/BinaryDataByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	// TabularDataByFilter queries tabular data and metadata based on given filters.
	TabularDataByFilter(context.Context, *TabularDataByFilterRequest) (*TabularDataByFilterResponse, error)
	// BinaryDataByFilter queries binary data and metadata based on given filters.
	BinaryDataByFilter(context.Context, *BinaryDataByFilterRequest) (*BinaryDataByFilterResponse, error)
	// BinaryDataByIDs queries binary data and metadata based on given IDs.
	BinaryDataByIDs(context.Context, *BinaryDataByIDsRequest) (*BinaryDataByIDsResponse, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) TabularDataByFilter(context.Context, *TabularDataByFilterRequest) (*TabularDataByFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TabularDataByFilter not implemented")
}
func (UnimplementedDataServiceServer) BinaryDataByFilter(context.Context, *BinaryDataByFilterRequest) (*BinaryDataByFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BinaryDataByFilter not implemented")
}
func (UnimplementedDataServiceServer) BinaryDataByIDs(context.Context, *BinaryDataByIDsRequest) (*BinaryDataByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BinaryDataByIDs not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_TabularDataByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TabularDataByFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).TabularDataByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.data.v1.DataService/TabularDataByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).TabularDataByFilter(ctx, req.(*TabularDataByFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_BinaryDataByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDataByFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).BinaryDataByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.data.v1.DataService/BinaryDataByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).BinaryDataByFilter(ctx, req.(*BinaryDataByFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_BinaryDataByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDataByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).BinaryDataByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.data.v1.DataService/BinaryDataByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).BinaryDataByIDs(ctx, req.(*BinaryDataByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.app.data.v1.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TabularDataByFilter",
			Handler:    _DataService_TabularDataByFilter_Handler,
		},
		{
			MethodName: "BinaryDataByFilter",
			Handler:    _DataService_BinaryDataByFilter_Handler,
		},
		{
			MethodName: "BinaryDataByIDs",
			Handler:    _DataService_BinaryDataByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/data/v1/data.proto",
}
