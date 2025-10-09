package services

import (
	// "errors"
	// "fmt"
	"errors"
	"fmt"

	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
)

func AddToCart(cart []models.CartItem) error {

	for i := range cart {
		if cart[i].Quantity < 1 || cart[i].BookID < 1{
			return fmt.Errorf("validation failed for cart Item %d: missing required fields", i)
		}
		book, err := repositories.GetBook(cart[i].BookID)
		if err != nil {
			return err
		}
		if cart[i].Quantity > book.Stock{
			return errors.New("the stock is less than the required amount")
		}
		cart[i].Price = book.Price

		println(cart[i].Quantity)
		println(book.Stock)
	}
	return repositories.AddToCart(cart)
}

func GetCartItems() ([]models.CartItem, error) {
	return repositories.GetCartItems()
}

func SaveCart(userID uint) error {
	// cartItems, err := GetCartItems()
	// if err != nil{
	// 	return err
	// }

	return nil
}