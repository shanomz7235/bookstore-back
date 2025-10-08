package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
)

func SetupUserRoutes(app *fiber.App) {


	user := app.Group("/user")
	user.Post("/register", handlers.Register)
	user.Get("/:id", handlers.GetUser)
	user.Get("/", handlers.GetUsers)
	user.Post("/login", handlers.LoginUser)
	// book.Get("/:id", handlers.GetUser)
	// book.Get("/:id", handlers.GetBook)
	// book.Put("/:id", handlers.UpdateBook)
	// book.Delete("/:id", handlers.DeleteBook)
}
