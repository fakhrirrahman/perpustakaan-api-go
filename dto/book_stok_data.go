package dto

type CreateBookStokData struct {
	BookId int `json:"book_id" `
	Codes[] string `json:"codes" `
}

type DeleteBookStokData struct {
	BookId int `json:"book_id" `
	Codes[] string `json:"codes" `
}