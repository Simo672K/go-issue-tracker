package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func MigrateDB() {
	fmt.Println("Applying migration...")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while opening sql file")
		return
	}
	wd = wd + "/internal/db/migrations/"
	connStr := "postgres://postgres:123456@localhost:5432/issue_tracker/"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	sqlFile, err := os.ReadFile(wd + "001_database.sql")
	if err != nil {
		fmt.Println("Error while opening sql file")
		return
	}
	sql := string(sqlFile)

	commands := strings.Split(sql, ";")
	for _, command := range commands {
		command = strings.TrimSpace(command) // Clean up leading/trailing whitespace
		if command == "" {
			continue // Skip empty commands
		}
		_, err := db.Exec(command)
		if err != nil {
			log.Fatalf("Error executing command: %s, error: %v", command, err)
		}
	}

	fmt.Println("SQL script executed successfully!")
}
