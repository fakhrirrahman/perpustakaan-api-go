package service

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
)

type bookStockService struct {
	bookRepository     domain.BookRepository
	bookStokRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) *bookStockService {
	return &bookStockService{
		bookRepository:     bookRepository,
		bookStokRepository: bookStockRepository,
	}
}

func (b bookStockService) Create(ctx context.Context, req dto.CreateBookStokData) error {
	books, err := b.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}
	if books.ID == "" {
		return domain.BookNotFound
	}

	stocks := make([]domain.BookStock, 0)
	for _, v := range req.Codes {
		stocks = append(stocks, domain.BookStock{
			Code:   v,
			BookID: req.BookId,
			Status: domain.BookStockAvailable,
		})
	}

	return b.bookStokRepository.Save(ctx, stocks)
}

func (b bookStockService) Delete(ctx context.Context, req dto.DeleteBookStokData) error {
	return b.bookStokRepository.DeleteByCodes(ctx, req.Codes)
}
