package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Unique(),
		field.Float("price"),
		field.JSON("pictures", []string{}),
		field.JSON("categories", []string{}),
		field.Int32("qty_available"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("last_updated").
			Default(time.Now),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
