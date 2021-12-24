package main

import "context"

type Facade struct {
	sh *StudentHandler
}

func NewFacade(sh *StudentHandler) *Facade {
	return &Facade{sh}
}

func (f *Facade) GetStudents(ctx context.Context) {
	f.sh.GetStudents(ctx)
}
