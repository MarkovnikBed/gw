package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var cfg = Config{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   os.Getenv("PG_DBNAME"),
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		SSLMode:  os.Getenv("PG_SSLMODE"),
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
	if db, err := sql.Open("postgres", dsn); err != nil {

		log.Fatalf("не получилось создать соединение с БД(postgres) :%v", err)
	} else if err = db.Ping(); err != nil {
		log.Fatalf("не получилось проверить соединение с БД(postgres) :%v", err)
	} else {
		return db
	}
	return nil
}
