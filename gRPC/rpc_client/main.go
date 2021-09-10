package main

import (
	"context"
	"flag"
	"github.com/Vitokz/Task/config"
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultConfigPath = "config/config_grpc.toml"
)

func main() {

	var log = logrus.New()

	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}
	var from = flag.Arg(0)

	var to =flag.Arg(1)

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	conn,err := grpc.Dial(":"+cfg.Application.Port,grpc.WithInsecure())

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


