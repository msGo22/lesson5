package gorm_driver

import (
	"gorm.io/gorm"
	"week5/internal/core/domains"
)

type StudentGormRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) (*StudentGormRepo, error) {
	if err := db.AutoMigrate(&domains.Student{}); err != nil {
		return nil, err
	}
	return &StudentGormRepo{
		db: db,
	}, nil
}

func (s *StudentGormRepo) Save(student *domains.Student) error {
	return s.db.Create(student).Error
}

func (s *StudentGormRepo) GetByID(ID uint) (*domains.Student, error) {
	var student domains.Student
	if err := s.db.Find(&student, ID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
