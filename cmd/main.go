package main

import (
	"net/http"

	"go-starter/api"
	"go-starter/separateRepos/config"
	"go-starter/separateRepos/graceful"
	"go-starter/separateRepos/logger"
	"go-starter/separateRepos/utl"
)

var env api.Env

func main() {
	logger.Init(api.Name)
	config.MustLoadYamlEnv("config.yml", api.EnvPrefix, &env)
	env.Print()

	graceful.Server(&http.Server{
		Addr:              utl.HostPort("", env.Port),
		Handler:           api.Router(),
		ReadHeaderTimeout: env.HeaderTimeout,
	})
}
