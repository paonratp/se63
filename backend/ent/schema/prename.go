package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Prename holds the schema definition for the Prename entity.
type Prename struct {
	ent.Schema
}

// Fields of the Prename.
func (Prename) Fields() []ent.Field {
	return []ent.Field{
		field.String("prefix").NotEmpty(),
	}
}

// Edges of the Prename.
func (Prename) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prename2doctorinfo", Doctorinfo.Type).StorageKey(edge.Column("prefix")),
		edge.To("patientrecord", Patientrecord.Type).StorageKey(edge.Column("prefix")),
	}
}
