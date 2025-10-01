package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
)

func SetupBookRoutes(app *fiber.App) {
	app.Get("/", handlers.Welcome)

	book := app.Group("/books")
	book.Post("/", handlers.CreateBook)
	book.Get("/", handlers.GetBooks)
	book.Get("/:id", handlers.GetBook)
	book.Put("/:id", handlers.UpdateBook)
	book.Delete("/:id", handlers.DeleteBook)
}
