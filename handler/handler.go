package handler

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Port string
	Name string
	Log *logrus.Logger
	Db  Db
}

type Db interface {
	SetValue(ctx context.Context, key string, value string) error
	GetValue(ctx context.Context, key string) (string, error)
}
