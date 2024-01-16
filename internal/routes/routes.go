package routes

import (
	"screening/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()

	//We can add middlewares to this app by invoking app.Use()
	api := app.Group("/")
	api.Post("createUser", handler.CreateUser)
	api.Post("updateUser", handler.UpdateUser)

	return app
}
