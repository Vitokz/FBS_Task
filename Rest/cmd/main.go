package main

import (
	"flag"
	"github.com/Vitokz/Task/Rest/server"
	"github.com/Vitokz/Task/config"
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	defaultConfigPath = "config/config_rest.toml"
)

func main() {
	var log = logrus.New()

	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	hndlr := handler.Handler{
		Name: cfg.Application.Name,
		Port: cfg.Application.Port,
		Log:    log,
		Db:     repository.DbConn(cfg.DateBase.Port),
	}

	router := echo.New()

	rest := server.Rest{
		Router:  router,
		Handler: &hndlr,
	}
	rest.Route()
	server.Start(router, hndlr.Port)
	defer server.Stop(router, time.Hour*2)
}


