package college

import "math"

const (
	DefaultPageSize = 10
	DefaultPage     = 1
)

type Paginator struct {
	Total     int  `json:"total"`
	Page      int  `json:"page"`
	Pages     int  `json:"pages"`
	PageCount int  `json:"count"`
	Prev      bool `json:"prev"`
	Next      bool `json:"next"`
}

type PaginatorOptions func(*Paginator)

func NewPaginator(options ...PaginatorOptions) *Paginator {
	pg := &Paginator{
		PageCount: DefaultPageSize,
		Page:      DefaultPage,
	}

	for _, opt := range options {
		opt(pg)
	}

	return pg
}

func WithPaginatorPage(page int) PaginatorOptions {
	return func(p *Paginator) {
		if page != 0 {
			p.Page = page
		}
	}
}

func WithPaginatorCount(count int) PaginatorOptions {
	return func(p *Paginator) {
		if count != 0 {
			p.PageCount = count
		}
	}
}

func (pg *Paginator) setPages(total, count int) {
	pg.Pages = int(math.Ceil(float64(total) / float64(count)))
}

func (pg *Paginator) setPrev(page int) {
	hasPrev := false

	if page > 1 {
		hasPrev = true
	}

	pg.Prev = hasPrev
}

func (pg *Paginator) setNext(page, pages int) {
	hasNext := false

	if pages >= page+1 {
		hasNext = true
	}

	pg.Next = hasNext
}

func (pg *Paginator) getOffset() int {
	return pg.PageCount * (pg.Page - 1)
}

func (pg *Paginator) build(total int) {
	pg.setPages(total, pg.PageCount)
	pg.setNext(pg.Page, pg.Pages)
	pg.setPrev(pg.Page)
}
