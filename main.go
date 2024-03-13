package main

import (
	"fmt"
	"ngx/adapter/logger"
	"ngx/core/util"
	"os"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env config: %s", err)
		os.Exit(1)
	}

	logger := logger.NewLogger(config)
	logger.Info().Msg("This is a ngx project ✨")
}
