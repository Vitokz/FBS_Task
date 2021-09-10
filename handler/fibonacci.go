package handler

import (
	"fmt"
	"github.com/Vitokz/Task/models"
	"strconv"
	"strings"
)

var lastAddedInDB = 1

func (h *Handler) Fibonacci(from , to int) (*models.Response,error) {
	result := new(models.Response)
	for ; from <= to; from++ {
		fromStr := strconv.Itoa(from)
		val, err := h.Db.GetValue(fromStr)
		if err != nil {
			fib := strconv.Itoa(h.calculateFibonacciNumber(from))
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,fib)
			if err := h.Db.SetValue(fromStr, fib); err != nil {
				return nil, err
			}
		} else {
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,val)
		}
	}
	result.Numbers=strings.TrimSuffix(result.Numbers,", ")
	return result,nil
}

func (h *Handler) calculateFibonacciNumber(n int) int {
	if n == 0 {
		return 0
	}

	i:=lastAddedInDB+1
	aStr,_ := h.Db.GetValue(strconv.Itoa(i-2))
	bStr,_ := h.Db.GetValue(strconv.Itoa(i-1))

	a,_:=strconv.Atoi(aStr)
	b,_:=strconv.Atoi(bStr)
	for i:=lastAddedInDB+1; i < n; i++ {
		с := a + b
		a = b
		b = с
	}

	return a + b
}