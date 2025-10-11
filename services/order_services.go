package services

import (
	"errors"
	"fmt"

	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
)

func Purchase(userID uint) error {

	cart, err := repositories.GetCartItems(userID)
	if err != nil {
		return err
	}
	if len(cart.Items) == 0 {
		return errors.New("cart is empty")
	}

	totalPrice := CalculatePrice(cart.Items)

	newOrder := models.Order{
		UserID: userID,
		Items:  convertItemsToOrderItems(cart.Items),
		Total:  totalPrice,
		Status: "paid",
	}
	if err := repositories.CreateOrder(&newOrder); err != nil {
		return err
	}

	if err := UpdateBookStock(cart.Items); err != nil {
		return err
	}

	cart.Status = "completed"
	if err := repositories.UpdateCartStatus(cart); err != nil {
		return err
	}

	return nil
}

func GetOrders(userID uint) ([]models.OrderResponse, error) {
	orders, err := repositories.GetOrders(userID)
	if err != nil {
		return nil, err
	}

	return convertOrdersToResponse(orders)
}

func CalculatePrice(items []models.Items) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func convertItemsToOrderItems(cartItems []models.Items) []models.OrderItem {
	orderItems := make([]models.OrderItem, len(cartItems))

	for i, item := range cartItems {
		orderItems[i] = models.OrderItem{
			BookID:   item.BookID,
			Quantity: item.Quantity,
			Price:    item.Price,
		}
	}
	return orderItems
}

func UpdateBookStock(items []models.Items) error {
	for _, item := range items {
		book, err := repositories.GetBookEntityByID(item.BookID)
		if err != nil {
			return nil
		}

		if book.Stock < item.Quantity {
			return fmt.Errorf("book %d stock is insufficient", item.BookID)
		}

		book.Stock -= item.Quantity

		if err := repositories.UpdateBookStock(book); err != nil {
			return err
		}
	}
	return nil
}

func convertOrdersToResponse(orders []models.Order) ([]models.OrderResponse, error) {
	responses := make([]models.OrderResponse, 0, len(orders))
	for _, order := range orders {
		items := make([]models.OrderItemResponse, 0, len(order.Items))
		for _, item := range order.Items {
			items = append(items, models.OrderItemResponse{
				BookID:   item.BookID,
				Quantity: item.Quantity,
				Price:    item.Price,
			})
		}

		responses = append(responses, models.OrderResponse{
			ID:     order.ID,
			UserID: order.UserID,
			Items:  items,
			Total:  order.Total,
			Status: order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		})
	}
	return responses, nil
}

func UpdateOrderStatus(orderID uint, newStatus *models.Order) error {
	if newStatus.Status != "shipping" && newStatus.Status != "shipped"{
		return errors.New("invalid status")
	}
	order, err := repositories.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	order.Status = newStatus.Status

	if err := repositories.UpdateOrderStatus(order); err != nil {
		return err
	}


	return nil
}

func GetAllOrders() ([]models.OrderResponse, error) {
	orders, err := repositories.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return convertOrdersToResponse(orders)
}