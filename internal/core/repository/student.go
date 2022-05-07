package repository

import "week5/internal/core/domains"

type StudentRepo interface {
	Save(student *domains.Student) error
	GetByID(ID uint) (*domains.Student, error)
}
