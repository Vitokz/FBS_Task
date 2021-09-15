package handler

import (
	"github.com/Vitokz/Task/models"
	"strconv"
)

var lastAddedInDB = 2

func (h *Handler) Fibonacci(from , to int) (*models.Response,error) {
	result := new(models.Response)
	for ; from <= to; from++ {
		val, err := h.Db.GetFibonacci(from)
		if err != nil {
			fib := h.calculateFibonacciNumber(from)
			result.Numbers = append(result.Numbers, models.Fibonacci{
				Index: from,
				Fibonacci: fib,
			})
			if err := h.Db.SetFibonacci(from, fib); err != nil {
				return nil, err
			}
			lastAddedInDB=from
		} else {
			vrem ,_ := strconv.Atoi(val)
			result.Numbers = append(result.Numbers, models.Fibonacci{
				Index: from,
				Fibonacci: vrem,
			})
		}
	}
	return result,nil
}

func (h *Handler) calculateFibonacciNumber(n int) int {

	i:=lastAddedInDB+1
	aStr,_ := h.Db.GetFibonacci(i-2)
	bStr,_ := h.Db.GetFibonacci(i-1)

	a,_:=strconv.Atoi(aStr)
	b,_:=strconv.Atoi(bStr)
	c:=0
	for ; i <= n; i++ {
		c = a + b
		a = b
		b = c
		h.Db.SetFibonacci(i,c)
	}

	return c
}