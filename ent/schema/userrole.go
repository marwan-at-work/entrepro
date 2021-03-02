package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("target_id"),
		field.String("target_type"),
		field.String("actor_type"),
		field.Int("actor_id"),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("user_roles").Unique(),
		edge.From("user", User.Type).Ref("user_roles").Unique(),
	}
}
