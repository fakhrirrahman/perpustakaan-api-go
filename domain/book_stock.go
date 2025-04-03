package domain

import (
	"context"
	"database/sql"
	"go-web-native/dto"
)

const (
	BookStockAvailable = "available"
	BookStockBorrowed  = "borrowed"
)

type BookStock struct {
	Code string `db:"code"`
	BookId string `db:"book_id"`
	Status string `db:"status"`
	BorrowerId sql.NullString `db:"borrower_id"`
	BorrowedAt sql.NullTime `db:"borrowed_at"`
}

type BookStockRepository interface {
	FindBookId(ctx context.Context, id string) ([]BookStock, error)
	FindByBookAndCode(ctx context.Context, id string, code string) error
	Save(ctx context.Context, data []BookStock) error
	Update(ctx context.Context, stock *BookStock) error
	DeleteByBookId(ctx context.Context, id string) error
	DeleteByCodes(ctx context.Context, codes []string) error
}


type BookStockService interface {
	Create(ctx context.Context, req dto.CreateBookStokData) error
	Delete(ctx context.Context, req dto.DeleteBookStokData) error
}