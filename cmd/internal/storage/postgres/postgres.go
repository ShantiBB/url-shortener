package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"url-shortener/cmd/internal/storage/queries"
)

type Storage struct {
	Db *sql.DB
}

func New(storageURL string) (*Storage, error) {
	const op = "storage.New"

	db, err := sql.Open("pgx", storageURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec(queries.CreateURLTable)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{Db: db}, nil
}
