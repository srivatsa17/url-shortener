package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB user is empty")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB password is empty")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("DB name is empty")
	}

	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully")
	RunMigrations()
}

func RunMigrations() {
	if DB == nil {
		log.Fatal("DB is not initialized")
	}

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id BIGINT PRIMARY KEY,
		long_url TEXT NOT NULL,
		short_code TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		clicks INT DEFAULT 0
	);
	
	CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
	CREATE INDEX IF NOT EXISTS idx_created_at ON urls(created_at);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations applied successfully")
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
