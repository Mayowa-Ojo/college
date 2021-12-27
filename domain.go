package college

import (
	"college/ent"
	"context"

	"github.com/google/uuid"
)

type Datastore struct {
	c *ent.Client
}

type GetStudentsVM struct {
	Page  int `form:"page"`
	Count int `form:"count"`
}

type CreateStudentVM struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type UpdateStudentDetailsVM struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type UpdateStudentDepartmentVM struct {
	DepartmentID string `json:"department_id"`
}

type CreateDepartmentVM struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Telephone string `json:"telephone"`
}

type GetClassesVM struct {
	Page  int `form:"page"`
	Count int `form:"count"`
}

type ClassRegistrationVM struct {
	Code string `form:"code"`
}

type StudentRepository interface {
	FindListStudents(ctx context.Context, pg *Paginator) ([]*ent.Student, *Paginator, error)
	FindStudentByEmail(ctx context.Context, email string) (*ent.Student, error)
	FindStudentByID(ctx context.Context, ID uuid.UUID) (*ent.Student, error)
	CreateStudent(ctx context.Context, s *Student) (*ent.Student, error)
	UpdateStudentDetails(ctx context.Context, s *Student, ID uuid.UUID) (*ent.Student, error)
	UpdateStudentDepartment(ctx context.Context, studentID, departmentID uuid.UUID) (*ent.Student, error)
	DeleteStudent(ctx context.Context, ID uuid.UUID) error
	UpdateStudentAddClass(ctx context.Context, sID uuid.UUID, class *ent.Class) (*ent.Student, error)
	UpdateStudentRemoveClass(ctx context.Context, sID uuid.UUID, classCode string) (*ent.Student, error)
}

type DepartmentRepository interface {
	FindListDepartments(ctx context.Context) ([]*ent.Department, error)
	FindListDepartmentsWithStudents(ctx context.Context) ([]*ent.Department, error)
	FindDepartmentByID(ctx context.Context, ID uuid.UUID) (*ent.Department, error)
	FindDepartmentByName(ctx context.Context, name string) (*ent.Department, error)
	CreateDepartment(ctx context.Context, d *Department) (*ent.Department, error)
}

type ClassRepository interface {
	FindListClasses(ctx context.Context, pg *Paginator) ([]*ent.Class, *Paginator, error)
	FindClassByCode(ctx context.Context, code string) (*ent.Class, error)
	// FindListClassesForStudent(ctx context.Context, pg *Paginator) ([]*ent.Class, *Paginator, error)
	// FindListClassesForDepartment(ctx context.Context, pg *Paginator) ([]*ent.Class, *Paginator, error)
}

type StaffRepository interface {
	FindListStaffs(ctx context.Context, pg *Paginator) ([]*ent.Staff, *Paginator, error)
	FindStaffByID(ctx context.Context, ID uuid.UUID) (*ent.Staff, error)
	// CreateStaff(ctx context.Context, s *Staff) (*ent.Staff, error)
}
