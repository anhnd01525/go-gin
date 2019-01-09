package model

import (
	"time"
)

type Cart struct {
	TableName struct{} `json:"table_name" sql:"shop.carts"`

	Id string `json:"id"`

	UserId string `json:"user_id"`

	Total float64 `json:"total"`

	Time time.Time `json:"time"`
}

type Bill struct {
	UserId string `json:"id"`

	Products []Products
}

type Products struct {
	ProductId string `json:"id"`

	Quantity int16 `json:"quantity"`
}

type GetPurchaseHistory struct {
	CartId string `json:"cart_id"`

	Username string `json:"username"`

	Total float64 `json:"total"`

	Time time.Time `json:"time"`

	Detail []Details `json:"detail"`
}

type Details struct {
	ProductName string `json:"product_name"`

	Manufacturer string `json:"manufacturer"`

	Price float64 `json:"price"`

	Quantity int16 `json:"quantity"`

	Total float64 `json:"total"`
}