package handler

import (
	"context"
	"fmt"
	"github.com/Vitokz/Task/functions"
	"github.com/Vitokz/Task/models"
	"strconv"
	"strings"
)

func (h *Handler) Fibonacci(from , to int,ctx context.Context) (*models.Response,error) {
	name:=h.Name
	switch name {
	case "grpc":

	case "rest":
	}
	result := new(models.Response)
	for ; from <= to; from++ {
		fromStr := strconv.Itoa(from)
		val, err := h.Db.GetValue(ctx, fromStr)
		if err != nil {
			fib := strconv.Itoa(functions.CalculateFibonacciNumber(from))
			if from==101 {
				h.Log.Info(fib)
			}
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,fib)
			if err := h.Db.SetValue(ctx, fromStr, fib); err != nil {
				return nil, err
			}
		} else {
			if from==101 {
				h.Log.Info(val)
			}
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,val)
		}
	}
	result.Numbers=strings.TrimSuffix(result.Numbers,", ")
	return result,nil
}
