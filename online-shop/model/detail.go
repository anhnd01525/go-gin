package model

type Detail struct {
	TableName struct{} `json:"table_name" sql:"shop.details"`

	CartId string `json:"id"`

	ProductId string `json:"id"`

	Quantity int16 `json:"quantity"`

	Total float64 `json:"total"`
}
