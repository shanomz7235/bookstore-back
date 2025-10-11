package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/services"
)

func Purchase(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID type",
		})
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := services.Purchase(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Purchase Complete",
	})
}

func GetOrderUser(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID type",
		})
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	orders, err := services.GetOrders(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(orders)

}

func UpdateOrderStatus(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}

	newStatus := new(models.Order)

	if err := c.BodyParser(newStatus); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Information",
		})
	}

	if err := services.UpdateOrderStatus(uint(orderID), newStatus); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Update Successful",
	})
}

func GetAllOrders(c *fiber.Ctx) error {
	orders, err := services.GetAllOrders()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(orders)
}