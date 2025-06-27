package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string         `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}