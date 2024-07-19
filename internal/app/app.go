package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"realworld-fiber-sqlc/internal/controller/http"
	"realworld-fiber-sqlc/pkg/logger"
	"realworld-fiber-sqlc/usecase/dto"
	"realworld-fiber-sqlc/usecase/dto/sqlc"
)

func New() *fiber.App {
	l := logger.New("debug")

	var err error
	pool, err := dto.NewPool(l)
	if err != nil {
		l.Fatal(err)
	}
	defer pool.Close()
	dbQueries := sqlc.New(pool)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	routes.Setup(app, dbQueries, l)

	return app
}
