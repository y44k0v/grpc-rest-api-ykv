package shop

import (
	"context"
	"log"

	product "github.com/y44k0v/grpc-rest-api-ykv/gen/product"
)

// product service example implementation.
// The example methods log the requests and return zero values.
type productsrvc struct {
	logger *log.Logger
}

// NewProduct returns the product service implementation.
func NewProduct(logger *log.Logger) product.Service {
	return &productsrvc{logger}
}

// Addproduct implements addproduct.
func (s *productsrvc) Addproduct(ctx context.Context, p *product.AddproductPayload) (err error) {
	s.logger.Print("product.addproduct")
	return
}

// Getproduct implements getproduct.
func (s *productsrvc) Getproduct(ctx context.Context, p *product.GetproductPayload) (res *product.ProductManagement, err error) {
	res = &product.ProductManagement{}
	s.logger.Print("product.getproduct")
	return
}

// Showproduct implements showproduct.
func (s *productsrvc) Showproduct(ctx context.Context) (res product.ProductManagementCollection, err error) {
	s.logger.Print("product.showproduct")
	return
}
