package main

import (
	"net/http"

	"go-starter/api"
	"go-starter/separateRepos/graceful"
	"go-starter/separateRepos/logger"
	"go-starter/separateRepos/utl"
)

func main() {
	logger.Init(ApiName)
	loadConfig("", "config.yml")

	graceful.Server(&http.Server{
		Addr:              utl.HostPort("", env.Port),
		Handler:           api.Router(),
		ReadHeaderTimeout: env.HeaderTimeout,
	})
}
