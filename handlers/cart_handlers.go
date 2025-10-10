package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/services"
)

func AddToCart(c *fiber.Ctx) error {
	items := []models.Items{}

	if err := c.BodyParser(&items); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Item Information",
		})
	}
	userIDStr, ok := c.Locals("user_id").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid user ID type",
        })
    }
	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err,
        })
	}
	if err := services.AddToCart(items, uint(userID)); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Message":"Add to cart successful",
		"Count": len(items),
	})
}

func GetCartItems(c *fiber.Ctx) error {
	items, err := services.GetCartItems()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}
	return c.JSON(items)
}

func SaveCart(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid user ID type",
        })
    }
	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err,
        })
	}

	if err := services.SaveCart(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}
	return c.JSON(fiber.Map{
		"Message": "save to cart successful",
		"user_ID": userID,
	})

}
