package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Pet.Type),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
