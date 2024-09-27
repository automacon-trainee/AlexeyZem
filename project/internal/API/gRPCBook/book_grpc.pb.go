// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/book.proto

package gRPCBook

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
	BookService_Take_FullMethodName   = "/gRPCBook.BookService/Take"
	BookService_Return_FullMethodName = "/gRPCBook.BookService/Return"
	BookService_Get_FullMethodName    = "/gRPCBook.BookService/Get"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Take(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error)
	Return(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Take(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Take_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Return(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Return_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Book, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility.
type BookServiceServer interface {
	Take(context.Context, *ID) (*Book, error)
	Return(context.Context, *ID) (*Book, error)
	Get(context.Context, *ID) (*Book, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServiceServer struct{}

func (UnimplementedBookServiceServer) Take(context.Context, *ID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Take not implemented")
}
func (UnimplementedBookServiceServer) Return(context.Context, *ID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}
func (UnimplementedBookServiceServer) Get(context.Context, *ID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}
func (UnimplementedBookServiceServer) testEmbeddedByValue()                     {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Take_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Take(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Take_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Take(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Return_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Return(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gRPCBook.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Take",
			Handler:    _BookService_Take_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _BookService_Return_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book.proto",
}
