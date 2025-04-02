package service

import (
	"context"
	"database/sql"
	"go-web-native/domain"
	"go-web-native/dto"
	"time"

	"github.com/google/uuid"
)

type CustomerService struct{ CustomerRepository domain.CustomerRepository }

func NewCustomerService(CustomerRepository domain.CustomerRepository) *CustomerService {
	return &CustomerService{CustomerRepository: CustomerRepository}
}

func (c CustomerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.CustomerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Name: v.Name,
			Code: v.Code,
		})
	}
	return customerData, nil

}

func (c CustomerService) Created(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:   uuid.New().String(),
		Name: req.Name,
		Code: req.Code,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	return c.CustomerRepository.Save(ctx, &customer)

}