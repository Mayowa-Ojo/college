package college

import "errors"

var (
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrDepartmentAlreadyExists = errors.New("department already exists")
)
