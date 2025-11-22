package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseUrl string) *sql.DB {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("Connected to DB")
	return db
}
