package dto

// CustomerData represents a customer entity
type CustomerData struct {
	ID   string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name string `json:"name" example:"John Doe"`
	Code string `json:"code" example:"CUST001"`
}

// CreateCustomerRequest represents the request to create a new customer
type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required" example:"John Doe" minLength:"2" maxLength:"255"`
	Code string `json:"code" validate:"required" example:"CUST001" minLength:"3" maxLength:"50"`
}

// UpdateCustomerRequest represents the request to update an existing customer
type UpdateCustomerRequest struct {
	ID   string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name string `json:"name" validate:"required" example:"John Doe" minLength:"2" maxLength:"255"`
	Code string `json:"code" validate:"required" example:"CUST001" minLength:"3" maxLength:"50"`
}

