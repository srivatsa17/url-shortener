package repository

import (
	"database/sql"
	"time"
)

type URLRepository interface {
	Create(id int64, longURL, code string) (time.Time, error)
	Get(code string) (string, error)
}

type urlRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) URLRepository {
	return &urlRepository{
		DB: db,
	}
}

func (u *urlRepository) Create(id int64, longURL, code string) (time.Time, error) {
	var createdAt time.Time
	err := u.DB.QueryRow(
		"INSERT INTO urls (id, long_url, short_code) VALUES ($1, $2, $3) RETURNING created_at",
		id, longURL, code,
	).Scan(&createdAt)

	return createdAt, err
}

func (u *urlRepository) Get(code string) (string, error) {
	var longURL string
	err := u.DB.QueryRow(
		"SELECT long_url FROM urls WHERE short_code=$1",
		code,
	).Scan(&longURL)

	return longURL, err
}
