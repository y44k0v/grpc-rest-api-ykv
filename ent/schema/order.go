package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.JSON("products", map[int]int{}),
		field.Time("last_updated").
			Default(time.Now),
		field.Enum("status").
			Values("placed", "processing", "finished").
			Default("placed"),
		field.Float32("sub_total"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
