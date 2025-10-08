package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/services"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Bookstore API!")
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Book Information",
		})
	}
	if err := services.CreateBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Create Book Complete!",
	})
}

func GetBooks(c *fiber.Ctx) error {
	books, err := services.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}
	book, err := services.GetBook(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}

	newBook := new(models.BookUpdate)

	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Book Information",
		})
	}

	if err = services.UpdateBook(uint(id), newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Update Successful",
	})

}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid book id",
		})
	}
	err = services.DeleteBook(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Delete Book Complete",
	})
}
