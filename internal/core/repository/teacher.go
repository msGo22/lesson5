package repository

import "github.com/msGo22/lesson5/internal/core/domains"

type TeacherRepo interface {
	Save(teacher *domains.Teacher) error
	GetByID(ID uint) (*domains.Teacher, error)
	Araba()
}
