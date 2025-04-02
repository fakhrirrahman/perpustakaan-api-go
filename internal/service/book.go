package service

import (
	"context"
	"database/sql"
	"errors"
	"go-web-native/domain"
	"go-web-native/dto"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBookService(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) *bookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

func (b bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []dto.BookData
	for _, v := range result {
		data = append(data, dto.BookData{
			Id: v.Id,
			Isbn: v.Isbn,
			Title: v.Title,
			Description: v.Description,

		})
	}
	return data, nil
}

func (b bookService) Show(ctx context.Context, id string) (dto.BookData, error) {
	data, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookData{}, err
	}
	if data.Id == "" {
		return dto.BookData{}, errors.New("data book tidak ditemukan")
	}
	return dto.BookData{
		Id: data.Id,
		Isbn: data.Isbn,
		Title: data.Title,
		Description: data.Description,
	}, nil
}

func (b bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id: uuid.NewString(),
		Isbn: req.Isbn,
		Title: req.Title,
		Description: req.Description,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	return b.bookRepository.Save(ctx, &book)
}

func (b bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}
	if persisted.Id == "" {
		return errors.New("data book tidak ditemukan")
	}
	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	return b.bookRepository.Update(ctx, &persisted)
}

func (b bookService) Delete(ctx context.Context, id string) error {
	exitst, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exitst.Id == "" {
		return errors.New("data book tidak ditemukan")
	}
	err = b.bookRepository.Delete(ctx, exitst.Id)
	if err != nil {
		return err
		
	}
	return b.bookStockRepository.DeleteByBookId(ctx, exitst.Id)
}
