package main

import (
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../databases/migrations",
	}
}
