package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uncleTen276/attendance.git/server/api/controllers"
)

func AuthRoutes(app fiber.Router) {
	app.Post("/register", controllers.CreateUser)
	app.Post("/login", controllers.Login)
}
