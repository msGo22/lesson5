package repository

import (
	"gorm.io/gorm"
	"week5/internal/core/domains"
)

type StudentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) (*StudentRepo, error) {
	if err := db.AutoMigrate(&domains.Student{}); err != nil {
		return nil, err
	}
	return &StudentRepo{
		db: db,
	}, nil
}

func (s *StudentRepo) Save(student *domains.Student) error {
	return s.db.Save(student).Error
}

func (s *StudentRepo) GetByID(ID uint) (*domains.Student, error) {
	var student domains.Student
	if err := s.db.Find(&student, ID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
