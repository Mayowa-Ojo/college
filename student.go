package college

import (
	"college/ent"
	"college/ent/student"
	"context"
	"time"

	"github.com/google/uuid"
	nanoid "github.com/matoous/go-nanoid/v2"
)

/**
Model
*/
type Student struct {
	ID              uuid.UUID   `json:"id"`
	Firstname       string      `json:"firstname"`
	Lastname        string      `json:"lastname"`
	Email           string      `json:"email"`
	AdmissionNumber string      `json:"admission_number"`
	Year            int         `json:"year"`
	Department      *Department `json:"department"`
	Classes         []*Class    `json:"classes"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func parseToStudent(model *ent.Student) *Student {
	var d *Department
	var c []*Class

	if model.Edges.Department != nil {
		d = parseToDepartment(model.Edges.Department)
	}

	if len(model.Edges.Classes) >= 1 {
		for _, v := range model.Edges.Classes {
			c = append(c, parseToClass(v))
		}
	}

	return &Student{
		ID:              model.ID,
		Firstname:       model.Firstname,
		Lastname:        model.Lastname,
		Email:           model.Email,
		AdmissionNumber: model.AdmissionNumber,
		Year:            model.Year,
		Department:      d,
		Classes:         c,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
	}
}

func NewStudentService(client *ent.Client) StudentRepository {
	return &Datastore{c: client}
}

/**
Repository
*/
func (d *Datastore) FindListStudents(ctx context.Context, pg *Paginator) ([]*ent.Student, *Paginator, error) {
	total, err := d.c.Student.Query().Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	pg.Total = total
	pg.build(total)

	offset := pg.getOffset()
	limit := pg.PageCount

	entities, err := d.c.Student.
		Query().
		WithDepartment().
		Order(ent.Desc(student.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	return entities, pg, nil
}

func (d *Datastore) FindStudentByID(ctx context.Context, ID uuid.UUID) (*ent.Student, error) {
	entity, err := d.c.Student.Query().
		Where(student.IDEQ(ID)).
		WithDepartment().
		WithClasses().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindStudentByEmail(ctx context.Context, email string) (*ent.Student, error) {
	entity, err := d.c.Student.Query().Where(student.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) CreateStudent(ctx context.Context, s *Student) (*ent.Student, error) {
	entity, err := d.c.Student.
		Create().
		SetFirstname(s.Firstname).
		SetLastname(s.Lastname).
		SetEmail(s.Email).
		SetAdmissionNumber(s.AdmissionNumber).
		SetYear(s.Year).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateStudentDetails(ctx context.Context, s *Student, ID uuid.UUID) (*ent.Student, error) {
	entity, err := d.c.Student.
		UpdateOneID(ID).
		SetFirstname(s.Firstname).
		SetLastname(s.Lastname).
		SetEmail(s.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateStudentDepartment(ctx context.Context, studentID, departmentID uuid.UUID) (*ent.Student, error) {
	entity, err := d.c.Student.
		UpdateOneID(studentID).
		SetDepartmentID(departmentID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateStudentAddClass(ctx context.Context, sID uuid.UUID, class *ent.Class) (*ent.Student, error) {
	entity, err := d.c.Student.
		UpdateOneID(sID).
		AddClasses(class).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	entity, err = d.c.Student.Query().
		Where(student.IDEQ(entity.ID)).
		WithDepartment().
		WithClasses().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateStudentRemoveClass(ctx context.Context, sID uuid.UUID, classCode string) (*ent.Student, error) {
	return nil, nil
}

func (d *Datastore) DeleteStudent(ctx context.Context, ID uuid.UUID) error {
	if err := d.c.Student.DeleteOneID(ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}

/**
Handlers
*/
type StudentHandler struct {
	sr StudentRepository
	cr ClassRepository
}

func NewStudentHandler(sr StudentRepository, cr ClassRepository) *StudentHandler {
	return &StudentHandler{sr, cr}
}

func (sh *StudentHandler) GetStudents(ctx context.Context, vm GetStudentsVM) ([]*Student, *Paginator, error) {
	paginator := NewPaginator(WithPaginatorCount(vm.Count), WithPaginatorPage(vm.Page))

	entities, paginator, err := sh.sr.FindListStudents(ctx, paginator)
	if err != nil {
		return nil, paginator, err
	}

	var students []*Student

	for _, v := range entities {
		students = append(students, parseToStudent(v))
	}

	return students, paginator, nil
}

func (sh *StudentHandler) GetStudentDetails(ctx context.Context, ID string) (*Student, error) {
	sID := uuid.MustParse(ID)

	entity, err := sh.sr.FindStudentByID(ctx, sID)
	if err != nil {
		return nil, err
	}

	return parseToStudent(entity), nil
}

func (sh *StudentHandler) CreateStudent(ctx context.Context, vm CreateStudentVM) (*Student, error) {
	_, err := sh.sr.FindStudentByEmail(ctx, vm.Email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	admissionNumber, err := nanoid.New(8)
	if err != nil {
		return nil, err
	}

	s := &Student{
		Firstname:       vm.Firstname,
		Lastname:        vm.Lastname,
		Email:           vm.Email,
		AdmissionNumber: admissionNumber,
		Year:            1,
	}

	entity, err := sh.sr.CreateStudent(ctx, s)
	if err != nil {
		return nil, err
	}

	return parseToStudent(entity), nil
}

func (sh *StudentHandler) UpdateStudentDetails(ctx context.Context, vm UpdateStudentDetailsVM, ID string) (*Student, error) {
	studentID := uuid.MustParse(ID)

	s := &Student{
		Firstname: vm.Firstname,
		Lastname:  vm.Lastname,
		Email:     vm.Email,
	}

	entity, err := sh.sr.UpdateStudentDetails(ctx, s, studentID)
	if err != nil {
		return nil, err
	}

	return parseToStudent(entity), nil
}

func (sh *StudentHandler) UpdateStudentDepartment(ctx context.Context, vm UpdateStudentDepartmentVM, ID string) (*Student, error) {
	studentID := uuid.MustParse(ID)
	departmentID := uuid.MustParse(vm.DepartmentID)

	entity, err := sh.sr.UpdateStudentDepartment(ctx, studentID, departmentID)
	if err != nil {
		return nil, err
	}

	return parseToStudent(entity), nil
}

func (sh *StudentHandler) ClassRegistration(ctx context.Context, vm ClassRegistrationVM, sID string) (*Student, error) {
	studentID := uuid.MustParse(sID)

	class, err := sh.cr.FindClassByCode(ctx, vm.Code)
	if err != nil {
		return nil, err
	}

	student, err := sh.sr.UpdateStudentAddClass(ctx, studentID, class)
	if err != nil {
		return nil, err
	}

	return parseToStudent(student), nil
}

func (sh *StudentHandler) DeleteStudent(ctx context.Context, ID string) error {
	studentID := uuid.MustParse(ID)

	if err := sh.sr.DeleteStudent(ctx, studentID); err != nil {
		return err
	}

	return nil
}
