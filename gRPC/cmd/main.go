package main

import (
	"github.com/Vitokz/Task/config"
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	configPath = "config/config_grpc.toml"
)

func main() {
	var log = logrus.New()

	cfg, err := config.Parse(configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	hndlr := handler.Handler{
		Port: cfg.Application.Port,
		Name: cfg.Application.Name,
		Log:    log,
		Db:     repository.DbConn(cfg.DateBase.Port),
	}

	{
	s:=grpc.NewServer()
	srv:=rpc.GRPCServer{Handler: &hndlr}
	rpc.RegisterFibonacciServer(s,&srv)

	lis,err := net.Listen("tcp",":"+hndlr.Port)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != err {
		log.Fatal(err)
	}
	}
}
