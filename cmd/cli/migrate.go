package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/ahmadkarlam/go-shipping/infrastructure/database"
)

func main() {
	migrations := &migrate.FileMigrationSource{
		Dir: "databases/migrations",
	}

	db, err := sql.Open("mysql", database.GetConnectionString())
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
