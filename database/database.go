package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("POSTGRESQL_HOST")
	username := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	databaseName := os.Getenv("POSTGRESQL_DATABASE")
	port := os.Getenv("POSTGRESQL_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
