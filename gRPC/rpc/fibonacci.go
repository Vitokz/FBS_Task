package rpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (g *GRPCServer) CalculateFibonacci(ctx context.Context, request *FibRequest) (*FibResponse, error) {
	from := request.GetFrom()
	if from == "" {
		err:=errors.New("Query param \"to\" is empty")
		g.Handler.Log.Error(err)
		return &FibResponse{},err
	}

	to := request.GetTo()
	if to == ""{
		err:=errors.New("Query param \"to\" is empty")
		g.Handler.Log.Error(err)
		return &FibResponse{},err
	}

	g.Handler.Log.WithFields(logrus.Fields{
		"event" : "Calculate Fibonacci",
		"from" : from,
		"to" : to,
	}).Info()

	resp,err := g.Handler.Fibonacci(from,to,ctx)
	if err != nil {
		g.Handler.Log.Error(err)
		return &FibResponse{},err
	}

	return &FibResponse{
		Numbers: resp.Numbers,
	},nil
}