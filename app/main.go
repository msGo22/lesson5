package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"week5/internal/core/domains"
	"week5/internal/core/repository"
	"week5/internal/core/services"
	"week5/pkg/repository/postgres"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Println(err)
		return
	}
	studentRepo, err := repository.NewStudentRepo(db)
	if err != nil {
		log.Println(err)
		return
	}
	studentService := services.NewStudentService(studentRepo)
	student1 := domains.NewStudent("ali")
	student2 := domains.NewStudent("veli")
	student3 := domains.NewStudent("cengiz")
	studentService.Create(student1)
	studentService.Create(student2)
	studentService.Create(student3)
	fmt.Println("-------")
	rstudent1, _ := studentService.GetByID(1)
	rstudent2, _ := studentService.GetByID(2)
	rstudent3, _ := studentService.GetByID(3)
	fmt.Println(rstudent1.Name)
	fmt.Println(rstudent2.Name)
	fmt.Println(rstudent3.Name)
}
