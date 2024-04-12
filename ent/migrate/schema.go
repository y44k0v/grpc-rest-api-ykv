// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "products_rodered", Type: field.TypeJSON},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"placed", "processing", "finished"}, Default: "placed"},
		{Name: "sub_total", Type: field.TypeFloat32},
		{Name: "user_orders", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_users_orders",
				Columns:    []*schema.Column{OrdersColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Unique: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "pictures", Type: field.TypeJSON},
		{Name: "categories", Type: field.TypeJSON},
		{Name: "qty_available", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "order_products", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_orders_products",
				Columns:    []*schema.Column{ProductsColumns[9]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString},
		{Name: "order_id", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"customer", "admin"}, Default: "customer"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
		ProductsTable,
		UsersTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
	ProductsTable.ForeignKeys[0].RefTable = OrdersTable
}
