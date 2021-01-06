package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Educationlevel holds the schema definition for the Educationlevel entity.
type Educationlevel struct {
	ent.Schema
}

// Fields of the Educationlevel.
func (Educationlevel) Fields() []ent.Field {
	return []ent.Field{
		field.String("level").NotEmpty(),
	}
}

// Edges of the Educationlevel.
func (Educationlevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("educationlevel2doctorinfo", Doctorinfo.Type).StorageKey(edge.Column("level")),
	}
}
