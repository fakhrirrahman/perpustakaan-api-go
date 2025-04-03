package service

import "go-web-native/domain"

type bookStockService struct {
	bookRepository domain.BookRepository
	bookStokRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, BookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &bookStockService{
		bookRepository:     bookRepository,
		bookStokRepository: bookStokRepository,
	}
}

func (b bookStockService) Create(ctx context.Context, req domain.CreateBookRequest) error {
	books, err := b.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}
	if books.Id == "" {
		return errors.New("data buku tidak ditemukan")
	}
	
	stocks := make([]domain.BookStock, 0)
	for i := 0; i < req.Stock; i++ {
		stocks = append(stocks, domain.BookStock{
			Code:     v,
			BookId:   req.BookId,
			Status:   domain.BookStockAvailable,
		})
	}
	return b.bookStokRepository.Save(ctx, stocks)

