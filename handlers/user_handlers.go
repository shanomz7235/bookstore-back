package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/services"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid User Formation",
		})
	}
	if err := services.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Register Successful!",
	})

}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	user, err := services.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(users)
}

func LoginUser(c *fiber.Ctx) error {
	user := new(models.LoginUser)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	t, err := services.LoginUser(user)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Unaothorized, Login fail",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})


	return c.JSON(fiber.Map{
		"message": "Login Successful!",
		// "token":t,
	})

}
