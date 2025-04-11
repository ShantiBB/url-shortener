package service

import (
	"fmt"
	"log/slog"
	"url-shortener/cmd/internal/repository"
)

type URLService struct {
	repo   *repository.Storage
	logger *slog.Logger
}

func NewURLService(repo *repository.Storage, log *slog.Logger) *URLService {
	return &URLService{
		repo:   repo,
		logger: log,
	}
}

func (s *URLService) SaveURL(url, alias string) error {
	id, err := s.repo.SaveURL(url, alias)
	if err != nil {
		s.logger.Error("failed to create URL", "alias", alias, "error", err)
		return fmt.Errorf("service.CreateURL: %w", err)
	}
	s.logger.Info("create URL", "id", id, "alias", alias)

	return nil
}

func (s *URLService) GetURL(alias string) error {
	url, err := s.repo.GetURL(alias)
	if err != nil {
		s.logger.Error("failed to retrieve URL", "error", err)
		return fmt.Errorf("service.RetrieveURL: %w", err)
	}
	s.logger.Info("get URL", "url", url)

	return nil
}

func (s *URLService) DeleteURL(alias string) error {
	err := s.repo.DeleteURL(alias)
	if err != nil {
		s.logger.Error("failed to remove URL", "error", err)
		return fmt.Errorf("service.RemoveURL: %w", err)
	}
	s.logger.Info("delete URL", "alias", alias)

	return nil
}
