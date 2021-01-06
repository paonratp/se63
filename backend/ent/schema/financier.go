package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)

// Financier holds the schema definition for the Financier entity.
type Financier struct {
	ent.Schema
}

// Fields of the Financier.
func (Financier) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Financier.
func (Financier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type).StorageKey(edge.Column("officer_id")),
		edge.From("user", User.Type).Ref("financier").Unique(),
	}
}
