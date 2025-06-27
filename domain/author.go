package domain

import (
	"context"
	"go-web-native/dto"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Bio       string         `json:"bio" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName untuk menentukan nama tabel di database
func (Author) TableName() string {
	return "authors"
}

type AuthorRepository interface {
	FindAll(ctx context.Context) ([]Author, error)
	FindById(ctx context.Context, id string) (Author, error)
	Save(ctx context.Context, author *Author) error
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id string) error
}

type AuthorService interface {
	Index(ctx context.Context) ([]dto.AuthorData, error)
	Show(ctx context.Context, id string) (dto.AuthorData, error)
	Create(ctx context.Context, req dto.CreateAuthorRequest) error
	Update(ctx context.Context, req dto.UpdateAuthorRequest) error
	Delete(ctx context.Context, id string) error
}
