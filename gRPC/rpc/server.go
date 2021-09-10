package rpc

import (
	"github.com/Vitokz/Task/handler"
)

type GRPCServer struct {
   Handler *handler.Handler
}


func (g *GRPCServer) mustEmbedUnimplementedFibonacciServer() {
    g.Handler.Log.Info()
}


