package service

import (
	"context"
	"errors"
	"go-web-native/domain"
	"go-web-native/dto"
	"go-web-native/internal/util"

	"github.com/google/uuid"
)

type authorService struct {
	authorRepo domain.AuthorRepository
}

func NewAuthorService(authorRepo domain.AuthorRepository) domain.AuthorService {
	return &authorService{
		authorRepo: authorRepo,
	}
}

func (s *authorService) Index(ctx context.Context) ([]dto.AuthorData, error) {
	authors, err := s.authorRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.AuthorData
	for _, author := range authors {
		result = append(result, dto.AuthorData{
			ID:        author.ID,
			Name:      author.Name,
			Email:     author.Email,
			Bio:       author.Bio,
			CreatedAt: author.CreatedAt,
			UpdatedAt: author.UpdatedAt,
		})
	}

	return result, nil
}

func (s *authorService) Show(ctx context.Context, id string) (dto.AuthorData, error) {
	author, err := s.authorRepo.FindById(ctx, id)
	if err != nil {
		return dto.AuthorData{}, err
	}

	result := dto.AuthorData{
		ID:        author.ID,
		Name:      author.Name,
		Email:     author.Email,
		Bio:       author.Bio,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}

	return result, nil
}

func (s *authorService) Create(ctx context.Context, req dto.CreateAuthorRequest) error {
	// Validasi request
	validationErrors := util.Validate(req)
	if len(validationErrors) > 0 {
		// Return first validation error as simple error
		for _, msg := range validationErrors {
			return errors.New(msg)
		}
	}

	// Generate UUID untuk ID
	authorID := uuid.New().String()

	author := &domain.Author{
		ID:    authorID,
		Name:  req.Name,
		Email: req.Email,
		Bio:   req.Bio,
	}

	return s.authorRepo.Save(ctx, author)
}

func (s *authorService) Update(ctx context.Context, req dto.UpdateAuthorRequest) error {
	// Validasi request
	validationErrors := util.Validate(req)
	if len(validationErrors) > 0 {
		// Return first validation error as simple error
		for _, msg := range validationErrors {
			return errors.New(msg)
		}
	}

	// Cek apakah author ada
	existingAuthor, err := s.authorRepo.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	// Update data
	existingAuthor.Name = req.Name
	existingAuthor.Email = req.Email
	existingAuthor.Bio = req.Bio

	return s.authorRepo.Update(ctx, &existingAuthor)
}

func (s *authorService) Delete(ctx context.Context, id string) error {
	return s.authorRepo.Delete(ctx, id)
}
