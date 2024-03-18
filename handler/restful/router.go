package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (server *Server) NewRouter() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.ConfigDefault))

	server.app = app
}
