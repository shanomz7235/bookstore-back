package models

type OrderItem struct {
    BookID   uint
    Quantity uint
    Price    uint
}

type Order struct {
    ID       uint
    UserID   uint
    Items    []OrderItem
    Total    uint
    Status   string // pending, paid, shipped
}
