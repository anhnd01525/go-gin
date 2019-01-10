package model

import "time"

type Detail struct {
	TableName struct {} `json:"table_name" sql:"library.details"`

	BorrowId string `json:"borrow_id"`

	BookId string `json:"book_id"`

	Returnday time.Time `json:"returnday"`

	Status int `json:"status"`
}

type ReturnBooks struct {
	BorrowId string `json:"borrow_id"`

	Books []Books
}

type Books struct {
	BookId string `json:"book_id"`

	Status string `json:"status"`
}