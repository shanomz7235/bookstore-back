package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/shanomz7235/bookstore-back/routes"
	"github.com/shanomz7235/bookstore-back/config"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	routes.SetupBookRoutes(app) 
	routes.SetupUserRoutes(app)
	routes.SetupCartRoutes(app)


	log.Println("Server running on :8080")
	app.Listen(":8080")
}
