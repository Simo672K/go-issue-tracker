package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

var PgStore *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error accured:", err)
		return
	}
	connStr := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error accured:", err)
		return
	}
	PgStore = db
}

func GetDBConn() *sql.DB {
	return PgStore
}
