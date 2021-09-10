package rpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (g *GRPCServer) CalculateFibonacci(ctx context.Context, request *FibRequest) (*FibResponse, error) {
	from := request.GetFrom()
	if from == "" {
		err := errors.New("Query param \"to\" is empty")
		g.Handler.Log.Error(err)
		return nil, err
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return nil, errors.New("\"from\" param is not number")
	} else if fromInt < 0 {
		return nil, errors.New("\"from\" param is minus")
	}

	to := request.GetTo()
	if to == "" {
		err := errors.New("Query param \"to\" is empty")
		g.Handler.Log.Error(err)
		return &FibResponse{}, err
	}
	toInt, err := strconv.Atoi(to)
	if err != nil {
		return nil, errors.New("\"to\" param is not number")
	} else if toInt < 0 {
		return nil, errors.New("\"to\" param is minus")
	} else if toInt < fromInt {
		return nil, errors.New("\"to\" param is less than \"from\"")
	} else if toInt > 92 {
		return nil, errors.New("\"to\" param must be less than or equal to 92")
	}

	g.Handler.Log.WithFields(logrus.Fields{
		"event": "Calculate Fibonacci",
		"from":  from,
		"to":    to,
	}).Info()

	resp, err := g.Handler.Fibonacci(fromInt, toInt)
	if err != nil {
		g.Handler.Log.Error(err)
		return &FibResponse{}, err
	}

	return &FibResponse{
		Numbers: resp.Numbers,
	}, nil
}
