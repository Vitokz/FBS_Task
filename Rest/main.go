package main

import (
	"github.com/Vitokz/Task/Rest/server"
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	port = "3000"
)

func main() {
	var log = logrus.New()

	hndlr := handler.Handler{
		Log:    log,
		Db:     repository.DbConn(),
	}

	router := echo.New()

	rest := server.Rest{
		Router:  router,
		Handler: &hndlr,
	}
	rest.Route()
	server.Start(router, port)
	defer server.Stop(router, time.Hour*2)
}


