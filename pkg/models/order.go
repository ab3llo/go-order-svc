package models

type Order struct {
	Id        string  `json:"id" gorm:"primaryKey"`
	Price     float32 `json:"price"`
	ProductId string  `json:"productId"`
	UserId    string  `json:"userId"`
}
