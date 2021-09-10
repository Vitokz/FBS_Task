package server

import (
	"context"
	"fmt"
	"github.com/Vitokz/Task/handler"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

const (
 apiPrefix = "/api"
)

type Rest struct {
	Router *echo.Echo
	Handler *handler.Handler
}

func (r *Rest) Route() {
	g := r.Router.Group(apiPrefix)
	g.GET("/fibonacci",r.Fibbonaci)
}

func Start(e *echo.Echo, addr string) {
	log:=logrus.New()
	port := fmt.Sprintf(":%v", addr)
	if err := e.Start(port); err != nil {
		log.Fatalf("[WARN] shutting down the server: %v", err)
	}
}

func Stop(e *echo.Echo, shutdownTimeout time.Duration) {
	log:=logrus.New()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error(" failed to shutdown the http server: %v", err)
		return
	}
	log.Info(" http server stopped")
}

