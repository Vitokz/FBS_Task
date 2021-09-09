package rpc

import (
   "github.com/Vitokz/Task/gRPC/handler"
)

type GRPCServer struct {
   Handler *handler.Handler
}


func (g *GRPCServer) mustEmbedUnimplementedFibonacciServer() {
    g.Handler.Log.Info()
}


