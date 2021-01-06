package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Treatment holds the schema definition for the Treatment entity.
type Treatment struct {
	ent.Schema
}

// Fields of the Treatment.
func (Treatment) Fields() []ent.Field {
	return []ent.Field{
		field.String("treatment").NotEmpty(),
		field.Time("datetime"),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return []ent.Edge{
	edge.From("typetreatment", Typetreatment.Type).Ref("treatment").Unique(),
	edge.From("patientrecord", Patientrecord.Type).Ref("treatment").Unique(),
	edge.From("doctorinfo", Doctorinfo.Type).Ref("treatment").Unique(),
	edge.To("unpaybills",Unpaybill.Type).StorageKey(edge.Column("treatment_id")).Unique(),
	}
}
