package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/harish876/go_api/configs"
	"github.com/harish876/go_api/router"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "https://github.com/harish876/go_api",
		})
	})
	api := app.Group("/api")
	router.UserRoutes(api.Group("/users"))
}
func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	configs.ConnectDB()
	setupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
