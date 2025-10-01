package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	fmt.Println("success connect migration")
	return goose.Up(db, "internal/infrastructure/persistence/postgresql/migrations")
}
