package dto

type CustomerData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}