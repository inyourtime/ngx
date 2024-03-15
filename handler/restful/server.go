package restful

import (
	"ngx/port"
	"ngx/util"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config  util.Config
	app     *fiber.App
	logger  port.Logger
	usecase port.Usecase
}

func NewServer(config util.Config, logger port.Logger, uc port.Usecase) port.Server {
	server := &Server{
		config:  config,
		logger:  logger,
		usecase: uc,
	}

	app := NewRouter()
	RegisterRoutes(app)

	server.app = app

	return server
}

func (server *Server) Start() error {
	return server.app.Listen(":" + server.config.ServerPort)
}
