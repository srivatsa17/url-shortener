package service

import (
	"fmt"

	"github.com/srivatsa17/url-shortener/model"
	"github.com/srivatsa17/url-shortener/repository"
	"github.com/srivatsa17/url-shortener/utils"
)

type URLService interface {
	ShortenURL(longURL string) (*model.URLResponse, error)
	GetURL(code string) (string, error)
}

type urlService struct {
	repo    repository.URLRepository
	baseURL string
}

func NewURLService(r repository.URLRepository, baseURL string) URLService {
	return &urlService{
		repo:    r,
		baseURL: baseURL,
	}
}

func (s *urlService) ShortenURL(longURL string) (*model.URLResponse, error) {
	// Validate URL format and protocol
	if err := utils.ValidateURL(longURL); err != nil {
		return nil, err
	}

	// Generate distributed ID
	id := utils.GenerateId()

	// Base62 encode the generated ID to get the short code
	code := utils.Encode(id)

	// Store the code in the db.
	createdAt, err := s.repo.Create(id, longURL, code)
	if err != nil {
		return nil, err
	}

	// Generate full short URL with domain prefix
	shortURL := fmt.Sprintf("%s/r/%s", s.baseURL, code)

	return &model.URLResponse{
		LongURL:   longURL,
		ShortURL:  shortURL,
		CreatedAt: createdAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *urlService) GetURL(code string) (string, error) {
	// Get the long URL
	longURL, err := s.repo.Get(code)
	if err != nil {
		return "", err
	}

	// Increment click count
	if err := s.repo.IncrementClickCount(code); err != nil {
		// Log error but don't fail the redirect
		fmt.Printf("failed to increment click count: %v\n", err)
	}

	return longURL, nil
}
