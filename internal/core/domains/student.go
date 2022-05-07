package domains

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	SchoolID uint   `json:"school_id"`
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}
