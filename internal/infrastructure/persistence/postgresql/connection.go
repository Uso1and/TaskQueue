package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {

	if err := godotenv.Load("configs/config.env"); err != nil {
		log.Println("warn: .env not loaded", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")
	if ssl == "" {
		ssl = "disable"
	}

	connect := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)

	var err error

	DB, err = sql.Open("postgres", connect)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Ping database success")
	return nil
}
