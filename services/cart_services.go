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
		items[i].CartID = cart.ID

	}
	return repositories.AddToCart(items)
}

func GetCartItems(id uint) (*models.CartResponse, error) {

	cart, err := repositories.GetCartItems(id)
	if err != nil{
		return nil, err
	}

	var items []models.ItemResponse
	for _, item := range cart.Items{
		items = append(items, models.ItemResponse{
			ID: item.ID,
			CartID: item.CartID,
			BookID: item.BookID,
			Quantity: item.Quantity,
			Price: item.Price,
		})
	}

	cartRes := &models.CartResponse{
		ID: cart.ID,
		UserID: cart.UserID,
		Status: cart.Status,
		Items: items,
	}
	return cartRes, nil
}


