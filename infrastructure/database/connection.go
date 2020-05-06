package database

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", GetConnectionString())
	if err != nil {
		panic(err)
	}

	return db
}

func DBTest() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	database, err := gorm.Open("mysql", db)
	if err != nil {
		panic(err)
	}

	return database, mock
}

func GetConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		username,
		password,
		hostname,
		port,
		database,
	)
}
