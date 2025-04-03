package dto

type CreateBookStokData struct {
	BookId string `json:"book_id" validate:"required"`
	Codes []string `json:"codes" validate:"required,unique,min=1"`
}

type DeleteBookStokData struct {
	Codes[] string `json:"codes" `
}