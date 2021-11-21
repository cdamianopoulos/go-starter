package main

import (
	"go-starter/api"
	"go-starter/separateRepos/config"
	"go-starter/separateRepos/graceful"
	"go-starter/separateRepos/logger"
	"go-starter/separateRepos/utl"
	"net/http"
)

func main() {
	// Initialise logging.
	logger.Init(api.Name)

	// Load environment config from config.yml and environment variables.
	var env api.Config
	config.MustLoadYamlEnv("config.yml", api.EnvPrefix, &env)
	env.Print()

	graceful.Server(&http.Server{
		Addr:              utl.HostPort("", env.Port),
		Handler:           api.Router(),
		ReadHeaderTimeout: env.HeaderTimeout,
	})
}
