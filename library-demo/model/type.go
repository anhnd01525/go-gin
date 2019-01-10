package model

type Type struct {
	TableName struct{} `json:"table_name" sql:"library.types"`

	Id string `json:"id"`

	Name string `json:"name"`

	BookCount int16 `json:"book_count"`
}

type CreateType struct {
	Name string `json:"name"`

	BookCount int16 `json:"book_count"`
}

type GetTypes struct {
	Id string `json:"id"`

	Name string `json:"name"`

	BookCount int16 `json:"book_count"`
}

type UpdateTypeById struct {
	Id string `json:"id"`

	Name string `json:"name"`

	BookCount int16 `json:"book_count"`
}