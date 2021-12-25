// Code generated by entc, DO NOT EDIT.

package ent

import (
	"college/ent/department"
	"college/ent/schema"
	"college/ent/student"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[1].Descriptor()
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	// departmentDescCode is the schema descriptor for code field.
	departmentDescCode := departmentFields[2].Descriptor()
	// department.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	department.CodeValidator = departmentDescCode.Validators[0].(func(string) error)
	// departmentDescTelephone is the schema descriptor for telephone field.
	departmentDescTelephone := departmentFields[3].Descriptor()
	// department.TelephoneValidator is a validator for the "telephone" field. It is called by the builders before save.
	department.TelephoneValidator = departmentDescTelephone.Validators[0].(func(string) error)
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentFields[4].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(func() time.Time)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentFields[5].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(func() time.Time)
	// department.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	department.UpdateDefaultUpdatedAt = departmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// departmentDescID is the schema descriptor for id field.
	departmentDescID := departmentFields[0].Descriptor()
	// department.DefaultID holds the default value on creation for the id field.
	department.DefaultID = departmentDescID.Default.(func() uuid.UUID)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescFirstname is the schema descriptor for firstname field.
	studentDescFirstname := studentFields[1].Descriptor()
	// student.FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	student.FirstnameValidator = studentDescFirstname.Validators[0].(func(string) error)
	// studentDescLastname is the schema descriptor for lastname field.
	studentDescLastname := studentFields[2].Descriptor()
	// student.LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	student.LastnameValidator = studentDescLastname.Validators[0].(func(string) error)
	// studentDescEmail is the schema descriptor for email field.
	studentDescEmail := studentFields[3].Descriptor()
	// student.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	student.EmailValidator = studentDescEmail.Validators[0].(func(string) error)
	// studentDescYear is the schema descriptor for year field.
	studentDescYear := studentFields[5].Descriptor()
	// student.YearValidator is a validator for the "year" field. It is called by the builders before save.
	student.YearValidator = func() func(int) error {
		validators := studentDescYear.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(year int) error {
			for _, fn := range fns {
				if err := fn(year); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescCreatedAt is the schema descriptor for created_at field.
	studentDescCreatedAt := studentFields[6].Descriptor()
	// student.DefaultCreatedAt holds the default value on creation for the created_at field.
	student.DefaultCreatedAt = studentDescCreatedAt.Default.(func() time.Time)
	// studentDescUpdatedAt is the schema descriptor for updated_at field.
	studentDescUpdatedAt := studentFields[7].Descriptor()
	// student.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	student.DefaultUpdatedAt = studentDescUpdatedAt.Default.(func() time.Time)
	// student.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	student.UpdateDefaultUpdatedAt = studentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentFields[0].Descriptor()
	// student.DefaultID holds the default value on creation for the id field.
	student.DefaultID = studentDescID.Default.(func() uuid.UUID)
}
