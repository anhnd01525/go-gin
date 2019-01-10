package model

type User struct {
	TableName struct{} `json:"table_name" sql:"library.users"`

	Id string `json:"id"`

	Fullname string `json:"fullname"`
}

type CreateUser struct {
	Fullname string `json:"fullname"`
}

type GetUsers struct {
	Id string `json:""id`

	Fullname string `json:"fullname"`
}
