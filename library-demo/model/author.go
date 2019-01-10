package model

type Author struct {
	TableName struct{} `json:"table_name" sql:"library.authors"`

	Id string `json:"id"`

	Fullname string `json:"fullname"`

	Address string `json:"address"`

	BookCount int16 `json:"book_count"`

	Introduction string `json:"introduction"`
}

type CreateAuthor struct {
	Fullname string `json:"fullname"`

	Address string `json:"address"`

	BookCount int16 `json:"book_count"`

	Introduction string `json:"introduction"`
}

type GetAuthors struct {
	Id string `json:"id"`

	Fullname string `json:"fullname"`

	Address string `json:"address"`

	BookCount int16 `json:"book_count"`

	Introduction string `json:"introduction"`
}

type UpdateAuthorById struct {
	Id string `json:"id"`

	Fullname string `json:"fullname"`

	Address string `json:"address"`

	BookCount int16 `json:"book_count"`

	Introduction string `json:"introduction"`
}
