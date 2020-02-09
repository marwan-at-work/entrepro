package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Spec holds the schema definition for the Spec entity.
type Spec struct {
	ent.Schema
}

// Fields of the Spec.
func (Spec) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(25).
			NotEmpty().
			Unique().
			Immutable(),
	}
}

// Edges of the Spec.
func (Spec) Edges() []ent.Edge {
	return nil
}
