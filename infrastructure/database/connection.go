package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", GetConnectionString())
	if err != nil {
		panic(err)
	}

	return db
}

func GetConnectionString() string {
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
