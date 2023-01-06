package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harish876/go_api/controllers"
)

func UserRoutes(route fiber.Router) {
	route.Get("/list", controllers.GetAllUsers)
	route.Get("/list/:id", controllers.GetUser)
	route.Post("/create", controllers.CreateUser)
	route.Patch("/update/:id", controllers.UpdateUser)
	route.Delete("/delete/:id", controllers.DeleteUser)
}
