package dto

type CreateBookStokData struct {
	BookId string `json:"book_id" validate:"required"`
	Codes []string `json:"codes" validate:"required,min=1,unique"`
}

type DeleteBookStokData struct {
	Codes[] string `json:"codes" `
}