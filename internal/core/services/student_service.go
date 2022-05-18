package services

import (
	"errors"
	"fmt"
	"github.com/msGo22/lesson5/internal/core/domains"
	"github.com/msGo22/lesson5/internal/core/repository"
)

type StudentService struct {
	studentRepo repository.StudentRepo
	teacherRepo repository.TeacherRepo
}

func NewStudentService(studentRepo repository.StudentRepo, teacherRepo repository.TeacherRepo) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
		teacherRepo: teacherRepo,
	}
}

func (s StudentService) Create(student *domains.Student) error {
	if student.Name == "" {
		return errors.New("student name can not be null")
	}
	return s.studentRepo.Save(student)
}

func (s StudentService) GetByID(ID uint) (*domains.Student, error) {
	return s.studentRepo.GetByID(ID)
}

func (s StudentService) CreateTeacher() error {
	t := domains.Teacher{
		Name:    "Ali Öğretmen",
		Section: "Beden Eğitimi",
	}
	fmt.Println(t.Name, " Eklendi")
	return s.teacherRepo.Save(&t)
}
