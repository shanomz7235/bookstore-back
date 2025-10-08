package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
	"github.com/shanomz7235/bookstore-back/middleware"
)

func SetupBookRoutes(app *fiber.App) {

	book := app.Group("/books")

	book.Use(middleware.AuthRequired)

	book.Get("/", handlers.GetBooks)
	book.Get("/:id", handlers.GetBook)

	admin := book.Group("/", middleware.RoleRequired("admin"))
	admin.Post("/", handlers.CreateBook)
	admin.Put("/:id", handlers.UpdateBook)
	admin.Delete("/:id", handlers.DeleteBook)
}
