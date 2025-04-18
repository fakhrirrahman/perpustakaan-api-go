package domain

import (
	"context"
	"database/sql"
	"go-web-native/dto"
	"time"
)

type Customer struct {
	ID        string    `db:"id"`
	Name      string `db:"name"`
	Code     string `db:"code"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"` 

}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
	Update(ctx context.Context, req dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, id string) error
	Show(ctx context.Context, id string) (dto.CustomerData, error)
}