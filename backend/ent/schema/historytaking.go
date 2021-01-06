package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Historytaking holds the schema definition for the Historytaking entity.
type Historytaking struct {
	ent.Schema
}

// Fields of the Historytaking.
func (Historytaking) Fields() []ent.Field {
	return []ent.Field{
		field.Float32("hight"),   //ส่วนสูง[เซนติเมตร]
		field.Float32("weight"),   //น้ำหนัก[กิโลกรัม]
		field.Float32("temp"),     //temperature ปกติ คือ 36.5-37.5 °C อุณภูมิร่างกาย[เซลเซียส]
		field.Int("pulse"),      //60-100 ครั้ง/นาที ชีพจร[ครั้งต่อนาที]
		field.Int("respiration"), //16-20ครั้งต่อนาที การหายใจ[ครั้งต่อนาที]
		field.Int("bp"),         //bloodpressure 80-120 มม.ปรอท  ความดันโลหิต[มิลลิเมตรปรอท]
		field.String("oxygen"),   //95 - 100% ออกซิเจนในเลือด[เปอร์เซ็นต์]
		field.String("symptom"),  //อาการ
		field.Time("datetime"),
	}
}

// Edges of the Historytaking.
func (Historytaking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nurse", Nurse.Type).Ref("historytaking").Unique(),
		edge.From("department", Department.Type).Ref("historytaking").Unique(),
		edge.From("symptomseverity", Symptomseverity.Type).Ref("historytaking").Unique(),
		edge.From("patientrecord", Patientrecord.Type).Ref("historytaking").Unique(),
	}
}
