package service

import (
	"context"
	"errors"
	"go-web-native/domain"
	"go-web-native/dto"

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
			Id: v.ID,
			Isbn: v.ISBN,
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
	if data.ID == "" {
		return dto.BookData{}, errors.New("data book tidak ditemukan")
	}
	return dto.BookData{
		Id: data.ID,
		Isbn: data.ISBN,
		Title: data.Title,
		Description: data.Description,
	}, nil
}

func (b bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		ID: uuid.NewString(),
		ISBN: req.Isbn,
		Title: req.Title,
		Description: req.Description,
	}
	return b.bookRepository.Save(ctx, &book)
}

func (b bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("data book tidak ditemukan")
	}
	persisted.ISBN = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	// UpdatedAt akan otomatis diupdate oleh GORM
	return b.bookRepository.Update(ctx, &persisted)
}

func (b bookService) Delete(ctx context.Context, id string) error {
	exitst, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exitst.ID == "" {
		return errors.New("data book tidak ditemukan")
	}
	err = b.bookRepository.Delete(ctx, exitst.ID)
	if err != nil {
		return err
		
	}
	return b.bookStockRepository.DeleteByBookId(ctx, exitst.ID)
}
