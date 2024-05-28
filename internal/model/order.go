package model

type Order struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}
