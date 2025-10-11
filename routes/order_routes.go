package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
	"github.com/shanomz7235/bookstore-back/middleware"
)

func SetupOrderRoutes(app *fiber.App) {

	order :=  app.Group("/order")

	order.Use(middleware.AuthRequired)

	order.Post("/purchase", handlers.Purchase)
	order.Get("/", handlers.GetOrderUser)
	admin := order.Group("/", middleware.RoleRequired("admin"))
	admin.Put("/:id", handlers.UpdateOrderStatus)
	admin.Get("/listOrder", handlers.GetAllOrders)
}