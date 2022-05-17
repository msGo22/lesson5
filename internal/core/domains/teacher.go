package domains

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string `json:"name"`
	Section string `json:"section"`
}

func NewTeacher(name, section string) *Teacher {
	return &Teacher{
		Name:    name,
		Section: section,
	}
}
