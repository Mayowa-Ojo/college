package college

import (
	"college/ent"
	"college/ent/class"
	"college/ent/schema"
	"context"
	"time"

	"github.com/google/uuid"
)

/**
Model
*/
type Class struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Code      string          `json:"code"`
	Unit      int             `json:"unit"`
	Location  string          `json:"location"`
	Semester  schema.Semester `json:"semester"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func parseToClass(model *ent.Class) *Class {
	return &Class{
		ID:        model.ID,
		Title:     model.Title,
		Code:      model.Code,
		Unit:      model.Unit,
		Location:  model.Location,
		Semester:  model.Semester,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func NewClassService(client *ent.Client) ClassRepository {
	return &Datastore{c: client}
}

/**
Repository
*/
func (d *Datastore) FindListClasses(ctx context.Context, pg *Paginator) ([]*ent.Class, *Paginator, error) {
	total, err := d.c.Class.Query().Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	pg.Total = total
	pg.build(total)

	offset := pg.getOffset()
	limit := pg.PageCount

	entities, err := d.c.Class.
		Query().
		Order(ent.Desc(class.FieldUnit)).
		Limit(limit).Offset(offset).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	return entities, pg, nil
}

func (d *Datastore) FindClassByCode(ctx context.Context, code string) (*ent.Class, error) {
	class, err := d.c.Class.Query().Where(class.CodeEQ(code)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return class, nil
}

/**
Handlers
*/
type ClassHandler struct {
	cr ClassRepository
}

func NewClassHandler(cr ClassRepository) *ClassHandler {
	return &ClassHandler{cr}
}

func (ch *ClassHandler) GetClasses(ctx context.Context, vm GetClassesVM) ([]*Class, *Paginator, error) {
	paginator := NewPaginator(WithPaginatorCount(vm.Count), WithPaginatorPage(vm.Page))

	entities, paginator, err := ch.cr.FindListClasses(ctx, paginator)
	if err != nil {
		return nil, nil, err
	}

	var classes []*Class

	for _, v := range entities {
		classes = append(classes, parseToClass(v))
	}

	return classes, paginator, nil
}
