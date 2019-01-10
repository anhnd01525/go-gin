package model

type Publisher struct {
	TableName struct{} `json:"table_name" sql:"library.publishers"`

	Id string `json:"id"`

	Name string `json:"name"`

	Address string `json:"address"`

	Fax string `json:"fax"`

	Mail string `json:"mail"`
}

type CreatePublisher struct {
	Name string `json:"name"`

	Address string `json:"address"`

	Fax string `json:"fax"`

	Mail string `json:"mail"`
}

type GetPublishers struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Address string `json:"address"`

	Fax string `json:"fax"`

	Mail string `json:"mail"`
}

type UpdatePublisherById struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Address string `json:"address"`

	Fax string `json:"fax"`

	Mail string `json:"mail"`
}