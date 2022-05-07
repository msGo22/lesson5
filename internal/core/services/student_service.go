package services

import (
	"week5/internal/core/domains"
	"week5/internal/core/repository"
)

type StudentService struct {
	studentRepo *repository.StudentRepo
}

func NewStudentService(studentRepo *repository.StudentRepo) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
	}
}

func (s StudentService) Create(student *domains.Student) error {
	return s.studentRepo.Save(student)
}

func (s StudentService) GetByID(ID uint) (*domains.Student, error) {
	return s.studentRepo.GetByID(ID)
}
