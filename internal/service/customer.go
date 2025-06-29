package service

import (
	"context"
	"errors"
	"go-web-native/domain"
	"go-web-native/dto"

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

func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:   uuid.New().String(),
		Name: req.Name,
		Code: req.Code,
		// CreatedAt akan otomatis diset oleh GORM
	}
	return c.CustomerRepository.Save(ctx, &customer)

}

func (c CustomerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.CustomerRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("data customer tidak ditemukan")
	}
	persisted.Name = req.Name
	persisted.Code = req.Code
	// UpdatedAt akan otomatis diupdate oleh GORM
	return c.CustomerRepository.Update(ctx, &persisted)
}

func (c CustomerService) Delete(ctx context.Context, id string) error {
	exitst, err := c.CustomerRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if exitst.ID == "" {
		return errors.New("data customer tidak ditemukan")
	}
	return c.CustomerRepository.Delete(ctx, id)
}

func (c CustomerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := c.CustomerRepository.FindByID(ctx, id)
	if err != nil {
		return dto.CustomerData{}, err
	}
	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("data customer tidak ditemukan")
	}
	return dto.CustomerData{
		ID:   persisted.ID,
		Name: persisted.Name,
		Code: persisted.Code,
	}, nil
}