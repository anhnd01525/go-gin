package model

type Book struct {
	TableName struct{} `json:"table_name" sql:"library.books"`

	Id string `json:"id"`

	Title string `json:title`

	AuthorId string `json:"author_id"`

	TypeId string `json:"type_id"`

	PublisherId string `json:"publisher_id"`

	Price float64 `json:"price"`

	Status int `json:"status" sql:",notnull"`
}

type CreateBook struct {
	Title string `json:title`

	AuthorId string `json:"author_id"`

	TypeId string `json:"type_id"`

	PublisherId string `json:"publisher_id"`

	Price float64 `json:"price"`

	Status int `json:"status" sql:",notnull"`
}

type GetBooks struct {
	Id string `json:"id"`

	Title string `json:title`

	AuthorId string `json:"author_id"`

	TypeId string `json:"type_id"`

	PublisherId string `json:"publisher_id"`

	Price float64 `json:"price"`

	Status int `json:"status" sql:",notnull"`
}

type UpdateBookById struct {
	Id string `json:"id"`

	Title string `json:title`

	AuthorId string `json:"author_id"`

	TypeId string `json:"type_id"`

	PublisherId string `json:"publisher_id"`

	Price float64 `json:"price"`

	Status int `json:"status" sql:",notnull"`
}

type GetBookById struct {
	Id string `json:"id"`

	Title string `json:title`

	AuthorId string `json:"author_id"`

	TypeId string `json:"type_id"`

	PublisherId string `json:"publisher_id"`

	Price float64 `json:"price"`

	Status int `json:"status" sql:",notnull"`
}