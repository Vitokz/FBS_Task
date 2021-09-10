package main

import (
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	configPath = "config/config_grpc.toml"
	port       = "3010"
)

func main() {
	var log = logrus.New()

	hndlr := handler.Handler{
		Log: log,
		Db:  repository.DbConn(),
	}

	s := grpc.NewServer()
	srv := rpc.GRPCServer{Handler: &hndlr}
	rpc.RegisterFibonacciServer(s, &srv)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != err {
		log.Fatal(err)
	}
}
