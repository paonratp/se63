package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Doctorinfo holds the schema definition for the Doctorinfo entity.
type Doctorinfo struct {
	ent.Schema
}

// Fields of the Doctorinfo.
func (Doctorinfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("doctorname").NotEmpty(),
		field.String("doctorsurname").NotEmpty(),
		field.String("telephonenumber").NotEmpty(),
		field.String("licensenumber").NotEmpty(),
	}
}

// Edges of the Doctorinfo.
func (Doctorinfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).Ref("department2doctorinfo").Unique(),
		edge.From("educationlevel", Educationlevel.Type).Ref("educationlevel2doctorinfo").Unique(),
		edge.From("officeroom", Officeroom.Type).Ref("officeroom2doctorinfo").Unique(),
		edge.From("prename", Prename.Type).Ref("prename2doctorinfo").Unique(),
		edge.From("user", User.Type).Ref("user2doctorinfo").Unique(),
		edge.To("treatment", Treatment.Type).StorageKey(edge.Column("doctorinfo_id")),
	}
}
