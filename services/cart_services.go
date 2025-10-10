package services

import (
	// "errors"
	// "fmt"
	"errors"
	"fmt"

	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
)

func AddToCart(items []models.Items, id uint) error {

	cart := repositories.GetCart(id)
	if cart == nil {
		cart = &models.Carts{UserID: id, Status: "active"}
		err := repositories.CreateCart(cart)
		if err != nil {
			return err
		}
	}

	for i := range items {
		if items[i].Quantity < 1 || items[i].BookID < 1 {
			// println(items[i].Quantity)
			// println(items[i].BookID)

			return fmt.Errorf("validation failed for cart Item %d: missing required fields", i)
		}
		book, err := repositories.GetBook(items[i].BookID)
		if err != nil {
			return err
		}
		if items[i].Quantity > book.Stock {
			return errors.New("the stock is less than the required amount")
		}
		items[i].Price = book.Price
		items[i].Cart_ID = cart.ID

	}
	return repositories.AddToCart(items)
}

func GetCartItems() ([]models.Items, error) {
	return repositories.GetItems()
}

func SaveCart(userID uint) error {
	cartItems, err := GetCartItems()
	if err != nil {
		return err
	}

	return repositories.SaveCart(cartItems, userID)
}
