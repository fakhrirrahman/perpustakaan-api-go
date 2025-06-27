package dto

import "time"

// AuthorData untuk response
type AuthorData struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name      string    `json:"name" example:"J.K. Rowling"`
	Email     string    `json:"email" example:"jk.rowling@example.com"`
	Bio       string    `json:"bio" example:"British author, best known for the Harry Potter series"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CreateAuthorRequest untuk membuat author baru
type CreateAuthorRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=255" example:"J.K. Rowling"`
	Email string `json:"email" validate:"required,email" example:"jk.rowling@example.com"`
	Bio   string `json:"bio" example:"British author, best known for the Harry Potter series"`
}

// UpdateAuthorRequest untuk update author
type UpdateAuthorRequest struct {
	ID    string `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name  string `json:"name" validate:"required,min=2,max=255" example:"J.K. Rowling"`
	Email string `json:"email" validate:"required,email" example:"jk.rowling@example.com"`
	Bio   string `json:"bio" example:"British author, best known for the Harry Potter series"`
}
