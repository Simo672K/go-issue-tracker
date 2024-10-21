package db

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

var connStr string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error accured:", err)
		return
	}
	connStr = os.Getenv("DATABASE")
}

func GetDBConnStr() string {
	return connStr
}
