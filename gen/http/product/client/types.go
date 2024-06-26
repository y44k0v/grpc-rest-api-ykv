// Code generated by goa v3.16.0, DO NOT EDIT.
//
// product HTTP client types
//
// Command:
// $ goa gen github.com/y44k0v/grpc-rest-api-ykv/design

package client

import (
	product "github.com/y44k0v/grpc-rest-api-ykv/gen/product"
	productviews "github.com/y44k0v/grpc-rest-api-ykv/gen/product/views"
	goa "goa.design/goa/v3/pkg"
)

// AddproductRequestBody is the type of the "product" service "addproduct"
// endpoint HTTP request body.
type AddproductRequestBody struct {
	// Product Name
	ProductName string `form:"ProductName" json:"ProductName" xml:"ProductName"`
	// Product description
	ProductDescription string `form:"ProductDescription" json:"ProductDescription" xml:"ProductDescription"`
	// Product Price
	ProductPrice float64 `form:"ProductPrice" json:"ProductPrice" xml:"ProductPrice"`
}

// GetproductResponseBody is the type of the "product" service "getproduct"
// endpoint HTTP response body.
type GetproductResponseBody struct {
	// SKU is the unique id of the Product.
	ProductSKU *string `form:"ProductSKU,omitempty" json:"ProductSKU,omitempty" xml:"ProductSKU,omitempty"`
	// Name of the Product.
	ProductName *string `form:"ProductName,omitempty" json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Description of the Product.
	ProductDescription *string `form:"ProductDescription,omitempty" json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	// Price of the Product.
	ProductPrice *float64 `form:"ProductPrice,omitempty" json:"ProductPrice,omitempty" xml:"ProductPrice,omitempty"`
}

// ShowproductResponseBody is the type of the "product" service "showproduct"
// endpoint HTTP response body.
type ShowproductResponseBody []*ProductManagementResponse

// ProductManagementResponse is used to define fields on response body types.
type ProductManagementResponse struct {
	// SKU is the unique id of the Product.
	ProductSKU *string `form:"ProductSKU,omitempty" json:"ProductSKU,omitempty" xml:"ProductSKU,omitempty"`
	// Name of the Product.
	ProductName *string `form:"ProductName,omitempty" json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Description of the Product.
	ProductDescription *string `form:"ProductDescription,omitempty" json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	// Price of the Product.
	ProductPrice *float64 `form:"ProductPrice,omitempty" json:"ProductPrice,omitempty" xml:"ProductPrice,omitempty"`
}

// NewAddproductRequestBody builds the HTTP request body from the payload of
// the "addproduct" endpoint of the "product" service.
func NewAddproductRequestBody(p *product.AddproductPayload) *AddproductRequestBody {
	body := &AddproductRequestBody{
		ProductName:        p.ProductName,
		ProductDescription: p.ProductDescription,
		ProductPrice:       p.ProductPrice,
	}
	return body
}

// NewGetproductProductManagementOK builds a "product" service "getproduct"
// endpoint result from a HTTP "OK" response.
func NewGetproductProductManagementOK(body *GetproductResponseBody) *productviews.ProductManagementView {
	v := &productviews.ProductManagementView{
		ProductSKU:         body.ProductSKU,
		ProductName:        body.ProductName,
		ProductDescription: body.ProductDescription,
		ProductPrice:       body.ProductPrice,
	}

	return v
}

// NewShowproductProductManagementCollectionOK builds a "product" service
// "showproduct" endpoint result from a HTTP "OK" response.
func NewShowproductProductManagementCollectionOK(body ShowproductResponseBody) productviews.ProductManagementCollectionView {
	v := make([]*productviews.ProductManagementView, len(body))
	for i, val := range body {
		v[i] = unmarshalProductManagementResponseToProductviewsProductManagementView(val)
	}

	return v
}

// ValidateProductManagementResponse runs the validations defined on
// ProductManagementResponse
func ValidateProductManagementResponse(body *ProductManagementResponse) (err error) {
	if body.ProductSKU == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductSKU", "body"))
	}
	if body.ProductName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductName", "body"))
	}
	if body.ProductDescription == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductDescription", "body"))
	}
	if body.ProductPrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductPrice", "body"))
	}
	return
}
