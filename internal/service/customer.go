package service

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
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