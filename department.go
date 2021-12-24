package main

import "ent-demo/ent"

func NewDepartmentService(client *ent.Client) DepartmentRepository {
	return &Datastore{c: client}
}

func (d *Datastore) FindListDepartments() {
	d.c.Student.Create()
}
