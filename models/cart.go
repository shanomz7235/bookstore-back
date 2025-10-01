package models

type CartItem struct {
    BookID   uint
    Quantity uint
    Price    uint
}

type Cart struct {
    UserID uint
    Items  []CartItem
}
