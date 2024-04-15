// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package productpb

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	// Addproduct implements addproduct.
	Addproduct(ctx context.Context, in *AddproductRequest, opts ...grpc.CallOption) (*AddproductResponse, error)
	// Getproduct implements getproduct.
	Getproduct(ctx context.Context, in *GetproductRequest, opts ...grpc.CallOption) (*GetproductResponse, error)
	// Showproduct implements showproduct.
	Showproduct(ctx context.Context, in *ShowproductRequest, opts ...grpc.CallOption) (*ProductManagementCollection, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Addproduct(ctx context.Context, in *AddproductRequest, opts ...grpc.CallOption) (*AddproductResponse, error) {
	out := new(AddproductResponse)
	err := c.cc.Invoke(ctx, "/product.Product/Addproduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Getproduct(ctx context.Context, in *GetproductRequest, opts ...grpc.CallOption) (*GetproductResponse, error) {
	out := new(GetproductResponse)
	err := c.cc.Invoke(ctx, "/product.Product/Getproduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Showproduct(ctx context.Context, in *ShowproductRequest, opts ...grpc.CallOption) (*ProductManagementCollection, error) {
	out := new(ProductManagementCollection)
	err := c.cc.Invoke(ctx, "/product.Product/Showproduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	// Addproduct implements addproduct.
	Addproduct(context.Context, *AddproductRequest) (*AddproductResponse, error)
	// Getproduct implements getproduct.
	Getproduct(context.Context, *GetproductRequest) (*GetproductResponse, error)
	// Showproduct implements showproduct.
	Showproduct(context.Context, *ShowproductRequest) (*ProductManagementCollection, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) Addproduct(context.Context, *AddproductRequest) (*AddproductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addproduct not implemented")
}
func (UnimplementedProductServer) Getproduct(context.Context, *GetproductRequest) (*GetproductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getproduct not implemented")
}
func (UnimplementedProductServer) Showproduct(context.Context, *ShowproductRequest) (*ProductManagementCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Showproduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_Addproduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddproductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Addproduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Addproduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Addproduct(ctx, req.(*AddproductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Getproduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetproductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Getproduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Getproduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Getproduct(ctx, req.(*GetproductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Showproduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowproductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Showproduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Showproduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Showproduct(ctx, req.(*ShowproductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addproduct",
			Handler:    _Product_Addproduct_Handler,
		},
		{
			MethodName: "Getproduct",
			Handler:    _Product_Getproduct_Handler,
		},
		{
			MethodName: "Showproduct",
			Handler:    _Product_Showproduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_grpc-rest-api-ykv_product.proto",
}
