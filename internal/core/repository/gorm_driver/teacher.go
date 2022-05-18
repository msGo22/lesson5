package gorm_driver

import (
	"fmt"
	"github.com/msGo22/lesson5/internal/core/domains"
	"gorm.io/gorm"
)

type TeacherGormDBRepo struct {
	db *gorm.DB
}

func NewTeacherGormDBRepo(db *gorm.DB) (*TeacherGormDBRepo, error) {
	if err := db.AutoMigrate(&domains.Teacher{}); err != nil {
		return nil, err
	}
	return &TeacherGormDBRepo{
		db: db,
	}, nil
}

func (t *TeacherGormDBRepo) Save(teacher *domains.Teacher) error {
	return t.db.Create(teacher).Error
}

func (t *TeacherGormDBRepo) GetByID(ID uint) (*domains.Teacher, error) {
	var teacher domains.Teacher
	if err := t.db.Find(&teacher, ID).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (t *TeacherGormDBRepo) Araba() {
	fmt.Println("Vın vın")
}
