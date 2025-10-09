package models

type CartItem struct {
	BookID   uint    `json:"bookid"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

type Cart struct {
	ID uint
	UserID uint
	Items  []CartItem
}
