package main

import (
	"flag"
	"github.com/Vitokz/Task/gRPC/config"
	"github.com/Vitokz/Task/gRPC/handler"
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/Vitokz/Task/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	defaultConfigPath = "gRPC/config/config.toml"
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

	{
	s:=grpc.NewServer()
	srv:=rpc.GRPCServer{Handler: &hndlr}
	rpc.RegisterFibonacciServer(s,&srv)

	lis,err := net.Listen("tcp",":"+cfg.Application.RpcPort)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != err {
		log.Fatal(err)
	}
	}
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
