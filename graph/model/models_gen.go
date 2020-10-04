// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	TotalPrice float64       `json:"totalPrice"`
	CartItems  []*CartOutput `json:"cartItems"`
}

type CartInput struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type CartOutput struct {
	Sku      string  `json:"sku"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
}

type CheckoutInput struct {
	TotalPrice float64      `json:"totalPrice"`
	CartItems  []*CartInput `json:"cartItems"`
}

type Item struct {
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderResult struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
