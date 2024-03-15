package main

import (
	"fmt"
	"log"
	"ngx/handler/restful"
	"ngx/logger"
	"ngx/util"
	"os"

	_ "ngx/docs"
)

// @title Ngx API
// @version 1.0
// @description This is a swagger for Ngx
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env config: %s", err)
		os.Exit(1)
	}

	logger := logger.NewLogger(config)
	logger.Info().Msg("This is a ngx project ✨")

	server := restful.NewServer(config, logger, nil)

	log.Fatal(server.Start())
}
