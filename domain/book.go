package domain

import (
	"context"
	"go-web-native/dto"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	ISBN        string         `json:"isbn" gorm:"type:varchar(20);unique;not null"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	Save(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}

type BookService interface {
	Index(ctx context.Context) ([]dto.BookData, error)
	Show(ctx context.Context, id string) (dto.BookData, error)
	Create(ctx context.Context, req dto.CreateBookRequest) error
	Update(ctx context.Context, req dto.UpdateBookRequest) error
	Delete(ctx context.Context, id string) error
}