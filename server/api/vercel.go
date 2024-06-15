package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/uncleTen276/attendance.git/server/api/routes"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("cannot load config", err)
	}

	if err := postgresql.ConnectDB(); err != nil {
		log.Fatal("cannot connect to db", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	corsApp := cors.ConfigDefault
	corsApp.AllowCredentials = true
	corsApp.AllowHeaders = "Origin, X-Requested-With, Content-Type, Accept, Authorization"
	app.Use(cors.New(corsApp))

	// app routes
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.AuthRoutes(api)

	return adaptor.FiberApp(app)
}
