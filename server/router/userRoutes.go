package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish876/go_api/controllers"
)

func UserRoutes(route fiber.Router) {
	route.Get("/list", controllers.GetAllUsers)
}
