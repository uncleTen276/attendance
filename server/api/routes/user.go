package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uncleTen276/attendance.git/server/api/controllers"
	middlewares "github.com/uncleTen276/attendance.git/server/api/middleware"
)

func UserRoutes(app fiber.Router) {
	user := app.Group("/user")

	privateUserRouter(user)
	publicUserRouter(user)
}

func publicUserRouter(app fiber.Router) {
}

func privateUserRouter(app fiber.Router) {
	app.Use(middlewares.JWTProtected())
	app.Get("/me", controllers.GetUserInfo)
}
