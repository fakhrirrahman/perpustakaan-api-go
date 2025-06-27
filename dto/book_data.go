package dto

// BookData represents a book entity
type BookData struct {
	Id          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Isbn        string `json:"isbn" example:"978-3-16-148410-0"`
	Title       string `json:"title" example:"Harry Potter and the Philosopher's Stone"`
	Description string `json:"description" example:"A young wizard's journey begins at Hogwarts School of Witchcraft and Wizardry"`
}

// CreateBookRequest represents the request to create a new book
type CreateBookRequest struct {
	Isbn        string `json:"isbn" validate:"required" example:"978-3-16-148410-0" minLength:"10" maxLength:"17"`
	Title       string `json:"title" validate:"required" example:"Harry Potter and the Philosopher's Stone" minLength:"1" maxLength:"255"`
	Description string `json:"description" validate:"required" example:"A young wizard's journey begins at Hogwarts School of Witchcraft and Wizardry" minLength:"1" maxLength:"1000"`
}

// UpdateBookRequest represents the request to update an existing book
type UpdateBookRequest struct {
	Id          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Isbn        string `json:"isbn" validate:"required" example:"978-3-16-148410-0" minLength:"10" maxLength:"17"`
	Title       string `json:"title" validate:"required" example:"Harry Potter and the Philosopher's Stone" minLength:"1" maxLength:"255"`
	Description string `json:"description" validate:"required" example:"A young wizard's journey begins at Hogwarts School of Witchcraft and Wizardry" minLength:"1" maxLength:"1000"`
}