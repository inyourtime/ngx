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

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
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
