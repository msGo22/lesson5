package services_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"week5/internal/core/domains"
	"week5/internal/core/services"
	"week5/pkg/mocks"
)

func TestNewStudentService(t *testing.T) {
	mockStudentRepo := new(mocks.MockStudentRepo)
	mockTeacherRepo := new(mocks.MockTeacherRepo)
	studentService := services.NewStudentService(mockStudentRepo, mockTeacherRepo)
	assert.NotNil(t, studentService)
}

func TestStudentService_Create(t *testing.T) {
	mockStudentRepo := new(mocks.MockStudentRepo)
	mockTeacherRepo := new(mocks.MockTeacherRepo)
	studentService := services.NewStudentService(mockStudentRepo, mockTeacherRepo)
	type input struct {
		studentName string
	}

	for _, test := range []struct {
		title string
		input
		want error
	}{
		{"success test", input{"test"}, nil},
		{"empty name", input{""}, errors.New("student name can not be null")},
	} {
		t.Run(test.title, func(t *testing.T) {
			student := domains.NewStudent(test.studentName)
			mockStudentRepo.On("Save", student).Return(nil)
			err := studentService.Create(student)
			assert.Equal(t, test.want, err)
		})
	}
}

func TestStudentService_GetByID(t *testing.T) {
	student := &domains.Student{Name: "test"}
	mockStudentRepo := new(mocks.MockStudentRepo)
	mockTeacherRepo := new(mocks.MockTeacherRepo)
	studentService := services.NewStudentService(mockStudentRepo, mockTeacherRepo)
	type input struct {
		id uint
	}
	type expected struct {
		student *domains.Student
		error   error
	}
	for _, test := range []struct {
		title string
		input
		expected
	}{
		{"success test", input{1}, expected{student: student}},
	} {
		t.Run(test.title, func(t *testing.T) {
			mockStudentRepo.On("GetByID", mock.Anything).Return(test.expected.student, test.expected.error)
			s, err := studentService.GetByID(test.input.id)
			assert.Equal(t, test.expected.student, s)
			assert.Equal(t, test.expected.error, err)
		})
	}
}
