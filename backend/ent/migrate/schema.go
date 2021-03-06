// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HistorytakingsColumns holds the columns for the "historytakings" table.
	HistorytakingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hight", Type: field.TypeFloat32},
		{Name: "weight", Type: field.TypeFloat32},
		{Name: "temp", Type: field.TypeFloat32},
		{Name: "pulse", Type: field.TypeInt},
		{Name: "respiration", Type: field.TypeInt},
		{Name: "bp", Type: field.TypeInt},
		{Name: "oxygen", Type: field.TypeString},
		{Name: "symptom", Type: field.TypeString},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "patientrecord_id", Type: field.TypeInt, Nullable: true},
		{Name: "symptomseverity_id", Type: field.TypeInt, Nullable: true},
	}
	// HistorytakingsTable holds the schema information for the "historytakings" table.
	HistorytakingsTable = &schema.Table{
		Name:       "historytakings",
		Columns:    HistorytakingsColumns,
		PrimaryKey: []*schema.Column{HistorytakingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "historytakings_departments_historytaking",
				Columns: []*schema.Column{HistorytakingsColumns[10]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_nurses_historytaking",
				Columns: []*schema.Column{HistorytakingsColumns[11]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_patientrecords_historytaking",
				Columns: []*schema.Column{HistorytakingsColumns[12]},

				RefColumns: []*schema.Column{PatientrecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_symptomseverities_historytaking",
				Columns: []*schema.Column{HistorytakingsColumns[13]},

				RefColumns: []*schema.Column{SymptomseveritiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NursesColumns holds the columns for the "nurses" table.
	NursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nursinglicense", Type: field.TypeString},
		{Name: "position", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// NursesTable holds the schema information for the "nurses" table.
	NursesTable = &schema.Table{
		Name:       "nurses",
		Columns:    NursesColumns,
		PrimaryKey: []*schema.Column{NursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "nurses_users_historytaking",
				Columns: []*schema.Column{NursesColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PatientrecordsColumns holds the columns for the "patientrecords" table.
	PatientrecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// PatientrecordsTable holds the schema information for the "patientrecords" table.
	PatientrecordsTable = &schema.Table{
		Name:        "patientrecords",
		Columns:     PatientrecordsColumns,
		PrimaryKey:  []*schema.Column{PatientrecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SymptomseveritiesColumns holds the columns for the "symptomseverities" table.
	SymptomseveritiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symptomseverity", Type: field.TypeString},
	}
	// SymptomseveritiesTable holds the schema information for the "symptomseverities" table.
	SymptomseveritiesTable = &schema.Table{
		Name:        "symptomseverities",
		Columns:     SymptomseveritiesColumns,
		PrimaryKey:  []*schema.Column{SymptomseveritiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
		HistorytakingsTable,
		NursesTable,
		PatientrecordsTable,
		SymptomseveritiesTable,
		UsersTable,
	}
)

func init() {
	HistorytakingsTable.ForeignKeys[0].RefTable = DepartmentsTable
	HistorytakingsTable.ForeignKeys[1].RefTable = NursesTable
	HistorytakingsTable.ForeignKeys[2].RefTable = PatientrecordsTable
	HistorytakingsTable.ForeignKeys[3].RefTable = SymptomseveritiesTable
	NursesTable.ForeignKeys[0].RefTable = UsersTable
}
