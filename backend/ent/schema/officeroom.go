package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Officeroom holds the schema definition for the Officeroom entity.
type Officeroom struct {
	ent.Schema
}

// Fields of the Officeroom.
func (Officeroom) Fields() []ent.Field {
	return []ent.Field{
		field.Int("roomnumber").Positive(),
	}
}

// Edges of the Officeroom.
func (Officeroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("officeroom2doctorinfo", Doctorinfo.Type).StorageKey(edge.Column("roomnumber")),
	}
}
