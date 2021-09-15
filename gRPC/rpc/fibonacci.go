package rpc

import (
	"context"
	"github.com/Vitokz/Task/models"
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
	} else if fromInt < 1 {
		return nil, errors.New("\"from\" param is minus or null")
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

	unResp, err := g.Handler.Fibonacci(fromInt, toInt)
	if err != nil {
		g.Handler.Log.Error(err)
		return &FibResponse{}, err
	}

	resp := sorting(unResp)

	return resp, nil
}

func sorting (resp *models.Response) *FibResponse {
	result := new(FibResponse)
	for i :=range resp.Numbers {
		result.Numbers = append(result.Numbers, &FibStruct{
              Index: int64(resp.Numbers[i].Index),
              Value: int64(resp.Numbers[i].Fibonacci),
		})
	}
	return result
}
