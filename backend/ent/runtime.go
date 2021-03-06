// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/schema"
	"github.com/team10/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescDepartment is the schema descriptor for department field.
	departmentDescDepartment := departmentFields[0].Descriptor()
	// department.DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	department.DepartmentValidator = departmentDescDepartment.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
