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
            "error": err.Error(),
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

	userIDStr, ok := c.Locals("user_id").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid user ID type",
        })
    }
	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
	}
	
	items, err := services.GetCartItems(uint(userID))
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(items)
}

func UpdateItems(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid user ID type",
        })
    }
	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
	}

	itemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}

	newItem := new(models.Items)

	if err := c.BodyParser(newItem); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Information",
		})
	}

	if err = services.UpdateItem(uint(userID), uint(itemID), newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Update Successful",
	})


}

func DeleteItem(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid user ID type",
        })
    }
	userID, err := strconv.Atoi(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
	}

	itemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}

	if err = services.DeleteItem(uint(userID), uint(itemID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Delete Item Successful",
	})
}