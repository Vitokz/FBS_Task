package main

import (
	"context"
	"flag"
	"github.com/Vitokz/Task/gRPC/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = "3010"
)

func main() {
	var log = logrus.New()
    flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}
	var from = flag.Arg(0)

	var to =flag.Arg(1)

	conn,err := grpc.Dial(":"+port,grpc.WithInsecure())

	c := rpc.NewFibonacciClient(conn)

	res,err :=c.CalculateFibonacci(context.Background(),&rpc.FibRequest{
		From: from,
		To: to,
	})
	if err !=nil {
		log.Fatal(err)
	}

	log.Println(res)
}


