package domain

import "errors"

var BookNotFound = errors.New("data buku tidak ditemukan")
var AuthorNotFound = errors.New("data penulis tidak ditemukan")
var AuthorEmailAlreadyExists = errors.New("email penulis sudah digunakan")