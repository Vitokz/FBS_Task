package handler

import (
	"context"
	"github.com/Vitokz/Task/Rest/config"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Config *config.Config
	Log    *logrus.Logger
	Db     Db
}

type Db interface {
	SetValue(ctx context.Context, key string, value string) error
	GetValue(ctx context.Context, key string) (string, error)
}
