// Code generated by goa v3.16.0, DO NOT EDIT.
//
// product gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/y44k0v/grpc-rest-api-ykv/design

package server

import (
	"context"

	productpb "github.com/y44k0v/grpc-rest-api-ykv/gen/grpc/product/pb"
	product "github.com/y44k0v/grpc-rest-api-ykv/gen/product"
	productviews "github.com/y44k0v/grpc-rest-api-ykv/gen/product/views"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAddproductResponse encodes responses from the "product" service
// "addproduct" endpoint.
func EncodeAddproductResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoAddproductResponse()
	return resp, nil
}

// DecodeAddproductRequest decodes requests sent to "product" service
// "addproduct" endpoint.
func DecodeAddproductRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *productpb.AddproductRequest
		ok      bool
	)
	{
		if message, ok = v.(*productpb.AddproductRequest); !ok {
			return nil, goagrpc.ErrInvalidType("product", "addproduct", "*productpb.AddproductRequest", v)
		}
	}
	var payload *product.AddproductPayload
	{
		payload = NewAddproductPayload(message)
	}
	return payload, nil
}

// EncodeGetproductResponse encodes responses from the "product" service
// "getproduct" endpoint.
func EncodeGetproductResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*productviews.ProductManagement)
	if !ok {
		return nil, goagrpc.ErrInvalidType("product", "getproduct", "*productviews.ProductManagement", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoGetproductResponse(result)
	return resp, nil
}

// DecodeGetproductRequest decodes requests sent to "product" service
// "getproduct" endpoint.
func DecodeGetproductRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *productpb.GetproductRequest
		ok      bool
	)
	{
		if message, ok = v.(*productpb.GetproductRequest); !ok {
			return nil, goagrpc.ErrInvalidType("product", "getproduct", "*productpb.GetproductRequest", v)
		}
	}
	var payload *product.GetproductPayload
	{
		payload = NewGetproductPayload(message)
	}
	return payload, nil
}

// EncodeShowproductResponse encodes responses from the "product" service
// "showproduct" endpoint.
func EncodeShowproductResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(productviews.ProductManagementCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("product", "showproduct", "productviews.ProductManagementCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoProductManagementCollection(result)
	return resp, nil
}
