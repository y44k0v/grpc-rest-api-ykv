// Code generated by goa v3.16.0, DO NOT EDIT.
//
// product gRPC server types
//
// Command:
// $ goa gen github.com/y44k0v/grpc-rest-api-ykv/design

package server

import (
	productpb "github.com/y44k0v/grpc-rest-api-ykv/gen/grpc/product/pb"
	product "github.com/y44k0v/grpc-rest-api-ykv/gen/product"
	productviews "github.com/y44k0v/grpc-rest-api-ykv/gen/product/views"
)

// NewAddproductPayload builds the payload of the "addproduct" endpoint of the
// "product" service from the gRPC request type.
func NewAddproductPayload(message *productpb.AddproductRequest) *product.AddproductPayload {
	v := &product.AddproductPayload{
		ProductSKU:         message.ProductSku,
		ProductName:        message.ProductName,
		ProductDescription: message.ProductDescription,
		ProductPrice:       message.ProductPrice,
	}
	return v
}

// NewProtoAddproductResponse builds the gRPC response type from the result of
// the "addproduct" endpoint of the "product" service.
func NewProtoAddproductResponse() *productpb.AddproductResponse {
	message := &productpb.AddproductResponse{}
	return message
}

// NewGetproductPayload builds the payload of the "getproduct" endpoint of the
// "product" service from the gRPC request type.
func NewGetproductPayload(message *productpb.GetproductRequest) *product.GetproductPayload {
	v := &product.GetproductPayload{
		ProductSKU: message.ProductSku,
	}
	return v
}

// NewProtoGetproductResponse builds the gRPC response type from the result of
// the "getproduct" endpoint of the "product" service.
func NewProtoGetproductResponse(result *productviews.ProductManagementView) *productpb.GetproductResponse {
	message := &productpb.GetproductResponse{
		ProductSku:         *result.ProductSKU,
		ProductName:        *result.ProductName,
		ProductDescription: *result.ProductDescription,
		ProductPrice:       *result.ProductPrice,
	}
	return message
}

// NewProtoProductManagementCollection builds the gRPC response type from the
// result of the "showproduct" endpoint of the "product" service.
func NewProtoProductManagementCollection(result productviews.ProductManagementCollectionView) *productpb.ProductManagementCollection {
	message := &productpb.ProductManagementCollection{}
	message.Field = make([]*productpb.ProductManagement, len(result))
	for i, val := range result {
		message.Field[i] = &productpb.ProductManagement{
			ProductSku:         *val.ProductSKU,
			ProductName:        *val.ProductName,
			ProductDescription: *val.ProductDescription,
			ProductPrice:       *val.ProductPrice,
		}
	}
	return message
}
