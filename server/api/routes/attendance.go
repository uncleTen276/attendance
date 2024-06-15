package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uncleTen276/attendance.git/server/api/controllers"
	middlewares "github.com/uncleTen276/attendance.git/server/api/middleware"
)

type AttendanceRouter struct {
	app fiber.Router
}

func NewAttendanceRouter(app fiber.Router) *AttendanceRouter {
	return &AttendanceRouter{
		app: app.Group("/attendance"),
	}
}

func (r *AttendanceRouter) RegisterRoutes() {
	r.privateRouters()
}

func (r *AttendanceRouter) privateRouters() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/", controllers.AttendanceRecordCheckIn)
	r.app.Get("/", controllers.GetAttendanceRecordByEmployee)
	r.app.Put("/", controllers.UpdateAttendance)
}
