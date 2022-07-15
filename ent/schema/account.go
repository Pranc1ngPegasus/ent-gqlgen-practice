package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("ID").
			Immutable().
			NotEmpty().
			Unique(),
		field.Time("created_at").
			Comment("Created at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Comment("Updated at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("email").
			Comment("Email").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
