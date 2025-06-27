package domain

import (
	"context"
	"go-web-native/dto"
	"time"

	"gorm.io/gorm"
)

const (
	BookStockAvailable = "available"
	BookStockBorrowed  = "borrowed"
)

type BookStock struct {
	ID         string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	Code       string         `json:"code" gorm:"type:varchar(50);unique;not null"`
	BookID     string         `json:"book_id" gorm:"type:varchar(36);not null"`
	Status     string         `json:"status" gorm:"type:varchar(20);not null;default:'available'"`
	BorrowerID *string        `json:"borrower_id" gorm:"type:varchar(36);default:null"`
	BorrowedAt *time.Time     `json:"borrowed_at" gorm:"type:timestamp;default:null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (BookStock) TableName() string {
	return "book_stocks"
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