// Code generated by goa v3.16.0, DO NOT EDIT.
//
// product views
//
// Command:
// $ goa gen github.com/y44k0v/grpc-rest-api-ykv/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ProductManagement is the viewed result type that is projected based on a
// view.
type ProductManagement struct {
	// Type to project
	Projected *ProductManagementView
	// View to render
	View string
}

// ProductManagementCollection is the viewed result type that is projected
// based on a view.
type ProductManagementCollection struct {
	// Type to project
	Projected ProductManagementCollectionView
	// View to render
	View string
}

// ProductManagementView is a type that runs validations on a projected type.
type ProductManagementView struct {
	// SKU is the unique id of the Product.
	ProductSKU *string
	// Name of the Product.
	ProductName *string
	// Description of the Product.
	ProductDescription *string
	// Price of the Product.
	ProductPrice *float64
}

// ProductManagementCollectionView is a type that runs validations on a
// projected type.
type ProductManagementCollectionView []*ProductManagementView

var (
	// ProductManagementMap is a map indexing the attribute names of
	// ProductManagement by view name.
	ProductManagementMap = map[string][]string{
		"default": {
			"ProductSKU",
			"ProductName",
			"ProductDescription",
			"ProductPrice",
		},
	}
	// ProductManagementCollectionMap is a map indexing the attribute names of
	// ProductManagementCollection by view name.
	ProductManagementCollectionMap = map[string][]string{
		"default": {
			"ProductSKU",
			"ProductName",
			"ProductDescription",
			"ProductPrice",
		},
	}
)

// ValidateProductManagement runs the validations defined on the viewed result
// type ProductManagement.
func ValidateProductManagement(result *ProductManagement) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProductManagementView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateProductManagementCollection runs the validations defined on the
// viewed result type ProductManagementCollection.
func ValidateProductManagementCollection(result ProductManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProductManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateProductManagementView runs the validations defined on
// ProductManagementView using the "default" view.
func ValidateProductManagementView(result *ProductManagementView) (err error) {
	if result.ProductSKU == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductSKU", "result"))
	}
	if result.ProductName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductName", "result"))
	}
	if result.ProductDescription == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductDescription", "result"))
	}
	if result.ProductPrice == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ProductPrice", "result"))
	}
	return
}

// ValidateProductManagementCollectionView runs the validations defined on
// ProductManagementCollectionView using the "default" view.
func ValidateProductManagementCollectionView(result ProductManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateProductManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
