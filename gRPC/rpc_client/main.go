package main

import (
	"context"
	"flag"
	"github.com/Vitokz/Task/gRPC/config"
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultConfigPath = "gRPC/config/config.toml"
	from="0"
	to="10"
)

func main() {

	var log = logrus.New()

	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	conn,err := grpc.Dial(":"+cfg.Application.RpcPort,grpc.WithInsecure())

	c := rpc.NewFibonacciClient(conn)

	res,err :=c.CalculateFibonacci(context.Background(),&rpc.FibRequest{
		From: from,
		To:to,
	})
	if err !=nil {
		log.Fatal(err)
	}
	log.Println(res)
}


