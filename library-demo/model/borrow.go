package model

import "time"

type Borrow struct {
	TableName struct{} `json:"table_name" sql:"library.borrows"`

	Id string `json:"id"`

	UserId string `json:"user_id"`

	Loanday time.Time `json:"loanday"`

	Returnterm time.Time `json:"returnterm"`
}

type Bill struct {
	UserId string `json:"user_id"`

	Books []string `json:"books"`
}

type Fines struct {
	Title string `json:"title"`

	OutOfDate int `json:"Số ngày quá hạn"`

	Money float64 `json:"Số tiền nộp phạt"`
}

type Lost struct {
	Title string `json:"title"`

	BookLost string `json:"Tình trạng"`

	Money float64 `json:"Số tiền nộp phạt"`
}


