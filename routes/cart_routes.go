package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
	"github.com/shanomz7235/bookstore-back/middleware"
)

func SetupCartRoutes(app *fiber.App) {

	cart := app.Group("/cart")

	cart.Use(middleware.AuthRequired)

	cart.Post("/", handlers.AddToCart)
	cart.Get("/", handlers.GetCartItems)

}