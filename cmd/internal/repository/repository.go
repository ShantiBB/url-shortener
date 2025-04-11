package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortener/cmd/internal/storage"
	"url-shortener/cmd/internal/storage/postgres"
	"url-shortener/cmd/internal/storage/queries"
)

type Storage struct {
	*postgres.Storage
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.SaveURL"

	stmt, err := s.Db.Prepare(queries.CreateURL)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id int64
	err = stmt.QueryRow(urlToSave, alias).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.GetURL"

	stmt, err := s.Db.Prepare(queries.GetURLByAlias)
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement %w", op, err)

	}

	var resURL string
	err = stmt.QueryRow(alias).Scan(&resURL)

	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound
	}
	if err != nil {
		return "", fmt.Errorf("%s: execute statement %w", op, err)
	}

	return resURL, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.DeleteURL"

	stmt, err := s.Db.Exec(queries.DeleteURLByAlias, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := stmt.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return storage.ErrURLNotFound
	}

	return nil
}
