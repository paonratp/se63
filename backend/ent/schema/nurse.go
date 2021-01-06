package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nursinglicense"),
		field.String("position"),
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("historytaking", Historytaking.Type).StorageKey(edge.Column("nurse_id")),
		edge.From("user", User.Type).Ref("historytaking").Unique(),
	}
}
