package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/uncleTen276/attendance.git/server/api/routes"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
)

func Serve() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("cannot load config", err)
	}

	if err := postgresql.ConnectDB(); err != nil {
		log.Fatal("cannot connect to db", err)
	}

	if err := postgresql.GetDB().Migrate(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.AuthRoutes(api)
	routes.NewAttendanceRouter(api).RegisterRoutes()
	routes.UserRoutes(api)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
