// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPictures holds the string denoting the pictures field in the database.
	FieldPictures = "pictures"
	// FieldCategories holds the string denoting the categories field in the database.
	FieldCategories = "categories"
	// FieldQtyAvailable holds the string denoting the qty_available field in the database.
	FieldQtyAvailable = "qty_available"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the product in the database.
	Table = "products"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "products"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPrice,
	FieldPictures,
	FieldCategories,
	FieldQtyAvailable,
	FieldCreatedAt,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_products",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastUpdated holds the default value on creation for the "last_updated" field.
	DefaultLastUpdated func() time.Time
)

// OrderOption defines the ordering options for the Product queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByQtyAvailable orders the results by the qty_available field.
func ByQtyAvailable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQtyAvailable, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLastUpdated orders the results by the last_updated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
	)
}
