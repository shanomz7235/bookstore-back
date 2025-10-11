package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shanomz7235/bookstore-back/handlers"
	"github.com/shanomz7235/bookstore-back/middleware"
)

func SetupUserRoutes(app *fiber.App) {

	user := app.Group("/user")

	user.Post("/register", handlers.Register)
	user.Post("/login", handlers.LoginUser)

	user.Use(middleware.AuthRequired)
	

	admin := user.Group("/", middleware.RoleRequired("admin"))
	admin.Get("/:id", handlers.GetUser)
	admin.Get("/", handlers.GetUsers)

}
