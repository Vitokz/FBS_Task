package handler

import (
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Port string
	Name string
	Log *logrus.Logger
	Db  Db
}

type Db interface {
	SetValue(key string, value string) error
	GetValue(key string) (string, error)
}
