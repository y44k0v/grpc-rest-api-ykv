/*
This is the design file. It contains the API specification, methods, inputs,
and outputs using Goa DSL code. The objective is to use this as a single
source of truth for the entire API source code.
*/
package design

import (
	. "goa.design/goa/v3/dsl"
)

// Main API declaration
var _ = API("shop", func() {
	Title("GRPC-REST api for users, orders, products")
	Description("This api manages CRUD operations")
	Server("shop", func() {
		Host("localhost", func() {
			URI("http://localhost:8080/api/v1")
			URI("grpc://localhost:9000")
		})
	})
})

// Order Service declaration with two methods and Swagger API specification file
var _ = Service("product", func() {
	Description("The Product service allows access to product members")
	Method("addproduct", func() {
		Payload(func() {
			Field(1, "ProductSKU", String, "Product SKU number")
			Field(2, "ProductName", String, "Product Name")
			Field(3, "ProductDescription", String, "Product description")
			Field(4, "ProductPrice", Float64, "Product Price")
			Required("ProductSKU", "ProductName", "ProductDescription", "ProductPrice")
		})
		Result(Empty)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			POST("/api/v1/prodcut/{ProductSKU}")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("getproduct", func() {
		Payload(func() {
			Field(1, "ProductSKU", String, "Product User")
			Required("ProductSKU")
		})
		Result(ProductManagement)
		Error("not_found", NotFound, "Product not found")
		HTTP(func() {
			GET("/api/v1/product/{ProductSKU}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("showproduct", func() {
		Result(CollectionOf(ProductManagement))
		HTTP(func() {
			GET("/api/v1/product")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Files("/openapi.json", "./gen/http/openapi.json")
})

// ProductManagement is a custom ResultType used to configure views for our custom type
var ProductManagement = ResultType("application/vnd.product", func() {
	Description("A ProductManagement type describes a product of shop.")
	Reference(Product)
	TypeName("ProductManagement")

	Attributes(func() {
		Field(1, "ProductSKU")
		Attribute("ProductSKU", String, "SKU is the unique id of the Product.", func() {
			Example("12356890")
		})
		Field(2, "ProductName")
		Attribute("ProductName", String, "Name of the Product.", func() {
			Example("Orange")
		})
		Field(3, "ProductDescription")
		Attribute("ProductDescription", String, "Description of the Product.", func() {
			Example("Organic orange of 234g")
		})
		Field(4, "ProductPrice")
		Attribute("ProductPrice", Float64, "Price of the Product.", func() {
			Example(6.85)
		})

	})

	View("default", func() {
		Field(1, "ProductSKU")
		Attribute("ProductSKU")
		Field(2, "ProductName")
		Attribute("ProductName")
		Field(3, "ProductDescription")
		Attribute("ProductDescription")
		Field(4, "ProductPrice")
		Attribute("ProductPrice")
	})

	Required("ProductSKU")
})

// User is the custom type for Users in the database
var Product = Type("Product", func() {
	Description("Product describes an item of shop.")
	Field(1, "ProductSKU")
	Attribute("ProductSKU", String, "SKU is the unique id of the Product.", func() {
		Example("12356890")
	})
	Field(2, "ProductName")
	Attribute("ProductName", String, "Name of the Product", func() {
		Example("Orange")
	})
	Field(3, "ProductDescription")
	Attribute("ProductDescription", String, "Description of the Product.", func() {
		Example("Organic orange of 234g")
	})
	Field(4, "ProductPrice")
	Attribute("ProductPrice", Float64, "Price of the Product.", func() {
		Example(6.85)
	})
	Required("ProductSKU", "ProductName", "ProductDescription", "ProductPrice")
})

// NotFound is a custom type where we add the queried field in the response
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when " +
		"the requested data that does not exist.")
	Field(1, "message")
	Attribute("message", String, "Message of error", func() {
		Example("Product 12356890 not found")
	})
	Field(2, "id", String, "ID of missing data")
	Required("message", "id")
})
