package repository

import "github.com/msGo22/lesson5/internal/core/domains"

type StudentRepo interface {
	Save(student *domains.Student) error
	GetByID(ID uint) (*domains.Student, error)
}
