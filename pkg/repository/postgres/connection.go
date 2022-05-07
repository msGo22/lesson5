package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connection() (*gorm.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	if username == "" || password == "" || dbName == "" {
		panic(fmt.Sprintln("DB Environments are missing", username, password, dbName))
	}
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", username, password, dbName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
