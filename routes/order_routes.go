package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
	"github.com/shanomz7235/bookstore-back/middleware"
)

func SetupOrderRoutes(app *fiber.App) {

	order :=  app.Group("/order")

	order.Use(middleware.AuthRequired)

	user := order.Group("/user", middleware.RoleRequired("user"))
	user.Post("/purchase", handlers.Purchase)
	user.Get("/", handlers.GetOrderUser)


	admin := order.Group("/admin", middleware.RoleRequired("admin"))
	admin.Put("/:id", handlers.UpdateOrderStatus)
	admin.Get("/listOrder", handlers.GetAllOrders)
}