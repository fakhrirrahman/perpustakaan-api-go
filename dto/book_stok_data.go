package dto

// CreateBookStokData represents the request to create book stock entries
type CreateBookStokData struct {
	BookId string   `json:"book_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Codes  []string `json:"codes" validate:"required,min=1,unique" example:"['BOOK001','BOOK002','BOOK003']"`
}

// DeleteBookStokData represents the request to delete book stock entries  
type DeleteBookStokData struct {
	Codes []string `json:"codes" example:"['BOOK001','BOOK002']"`
}

// BookStockData represents a book stock entry
type BookStockData struct {
	ID     string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	BookId string `json:"book_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Code   string `json:"code" example:"BOOK001"`
	Status string `json:"status" example:"available"`
}