package main

import (
	"flag"
	"github.com/Vitokz/Task/Rest/config"
	"github.com/Vitokz/Task/Rest/handler"
	"github.com/Vitokz/Task/Rest/server"
	"github.com/Vitokz/Task/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	defaultConfigPath = "Rest/config/config.toml"
)

func main() {
	var log = logrus.New()

	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	err = configureLogrus(log, cfg)
	if err != nil {
		log.Fatalf("failed to configure logrus: %v", err)
	}

	hndlr := handler.Handler{
		Config: cfg,
		Log:    log,
		Db:     repository.DbConn(cfg.DateBase.Port),
	}
	router := echo.New()

	rest := server.Rest{
		Router:  router,
		Handler: &hndlr,
	}
	rest.Route()
	server.Start(router, hndlr.Config.Application.HttpPort)
	defer server.Stop(router, time.Hour*2)
}

func configureLogrus(logger *logrus.Logger, cfg *config.Config) error {
	lvl, err := logrus.ParseLevel(cfg.Application.LogLevel)
	if err != nil {
		return err
	}
	logger.SetLevel(lvl)
	if cfg.Application.LogLevel == "text" {
		text := logrus.TextFormatter{}
		logger.SetFormatter(&text)
	}
	return nil
}
