package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patientrecord holds the schema definition for the Patientrecord entity.
type Patientrecord struct {
	ent.Schema
}

// Fields of the Patientrecord.
func (Patientrecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.Int("Idcardnumber"),
		field.Int("Age"),
		field.Time("Birthday"),
		field.String("Bloodtype"),
		field.String("Disease"),
		field.String("Allergic"),
		field.Int("Phonenumber"),
		field.String("Email"),
		field.String("Home"),
		field.Time("Date"),
	}
}

// Edges of the Patientrecord.
func (Patientrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("patientrecord").Unique(),
		edge.From("medicalrecordstaff", Medicalrecordstaff.Type).Ref("patientrecord").Unique(),
		edge.From("prename", Prename.Type).Ref("patientrecord").Unique(),
		edge.To("historytaking", Historytaking.Type).StorageKey(edge.Column("patientrecord_id")),
		edge.To("treatment", Treatment.Type).StorageKey(edge.Column("patientrecord_id")),
		edge.To("PatientrecordPatientrights", Patientrights.Type).StorageKey(edge.Column("patientrecord_id")),
	}
}
