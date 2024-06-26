// Code generated by goa v3.16.0, DO NOT EDIT.
//
// product HTTP client CLI support package
//
// Command:
// $ goa gen github.com/y44k0v/grpc-rest-api-ykv/design

package client

import (
	"encoding/json"
	"fmt"

	product "github.com/y44k0v/grpc-rest-api-ykv/gen/product"
)

// BuildAddproductPayload builds the payload for the product addproduct
// endpoint from CLI flags.
func BuildAddproductPayload(productAddproductBody string, productAddproductProductSKU string) (*product.AddproductPayload, error) {
	var err error
	var body AddproductRequestBody
	{
		err = json.Unmarshal([]byte(productAddproductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ProductDescription\": \"Voluptatem non quae vero eum non.\",\n      \"ProductName\": \"Nostrum aut necessitatibus.\",\n      \"ProductPrice\": 0.8245373747073658\n   }'")
		}
	}
	var productSKU string
	{
		productSKU = productAddproductProductSKU
	}
	v := &product.AddproductPayload{
		ProductName:        body.ProductName,
		ProductDescription: body.ProductDescription,
		ProductPrice:       body.ProductPrice,
	}
	v.ProductSKU = productSKU

	return v, nil
}

// BuildGetproductPayload builds the payload for the product getproduct
// endpoint from CLI flags.
func BuildGetproductPayload(productGetproductProductSKU string) (*product.GetproductPayload, error) {
	var productSKU string
	{
		productSKU = productGetproductProductSKU
	}
	v := &product.GetproductPayload{}
	v.ProductSKU = productSKU

	return v, nil
}
