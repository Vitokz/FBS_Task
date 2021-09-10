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
	SetFibonacci(key int,value int) error
	GetFibonacci(key int) (string,error)
}
