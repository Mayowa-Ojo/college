package college

import (
	"college/ent"
	"college/ent/schema"
	"college/ent/staff"
	"context"
	"time"

	"github.com/google/uuid"
)

/**
Model
*/
type Staff struct {
	ID        uuid.UUID        `json:"id"`
	Firstname string           `json:"firstname"`
	Lastname  string           `json:"lastname"`
	Email     string           `json:"email"`
	Telephone string           `json:"telephone"`
	Rank      schema.StaffRank `json:"rank"`
	Salary    int              `json:"salary"`
	Role      schema.StaffRole `json:"role"` // academic / non-academic
	Classes   []*Class         `json:"classes"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func parseToStaff(model *ent.Staff) *Staff {
	var c []*Class

	if len(model.Edges.Classes) >= 1 {
		for _, v := range model.Edges.Classes {
			c = append(c, parseToClass(v))
		}
	}

	return &Staff{
		ID:        model.ID,
		Firstname: model.Firstname,
		Lastname:  model.Lastname,
		Email:     model.Email,
		Telephone: model.Telephone,
		Rank:      model.Rank,
		Salary:    model.Salary,
		Role:      model.Role,
		Classes:   c,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func NewStaffService(client *ent.Client) StaffRepository {
	return &Datastore{c: client}
}

/**
Repository
*/
func (d *Datastore) FindListStaffs(ctx context.Context, pg *Paginator) ([]*ent.Staff, *Paginator, error) {
	total, err := d.c.Staff.Query().Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	pg.Total = total
	pg.build(total)

	offset := pg.getOffset()
	limit := pg.PageCount

	entities, err := d.c.Staff.
		Query().
		WithClasses().
		Order(ent.Desc(staff.FieldRank)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	return entities, nil, nil
}

func (d *Datastore) FindStaffByID(ctx context.Context, ID uuid.UUID) (*ent.Staff, error) {
	entity, err := d.c.Staff.Query().
		Where(staff.IDEQ(ID)).
		WithDepartment().
		WithClasses().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

/**
Handlers
*/
type StaffHandler struct {
	sr StaffRepository
	cr ClassRepository
}

func NewStaffHandler(sr StaffRepository, cr ClassRepository) *StaffHandler {
	return &StaffHandler{sr, cr}
}
