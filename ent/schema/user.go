package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").
			Unique(),
		field.String("address"),
		field.JSON("Order_id", []int{}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("last_updated").
			Default(time.Now),
		field.Enum("level").
			Values("customer", "admin").
			Default("customer"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
