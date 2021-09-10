package handler

import (
	"fmt"
	"github.com/Vitokz/Task/models"
	"strconv"
	"strings"
)

var lastAddedInDB = 2

func (h *Handler) Fibonacci(from , to int) (*models.Response,error) {
	result := new(models.Response)
	for ; from <= to; from++ {
		fromStr := strconv.Itoa(from)
		val, err := h.Db.GetFibonacci(from)
		if err != nil {
			fib := h.calculateFibonacciNumber(from)
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,strconv.Itoa(fib))
			if err := h.Db.SetFibonacci(from, fib); err != nil {
				return nil, err
			}
			lastAddedInDB=from
		} else {
			result.Numbers += fmt.Sprintf("[%s] = %s, ",fromStr,val)
		}
	}
	result.Numbers=strings.TrimSuffix(result.Numbers,", ")
	return result,nil
}

func (h *Handler) calculateFibonacciNumber(n int) int {

	i:=lastAddedInDB+1
	aStr,_ := h.Db.GetFibonacci(i-2)
	bStr,_ := h.Db.GetFibonacci(i-1)

	a,_:=strconv.Atoi(aStr)
	b,_:=strconv.Atoi(bStr)
	c:=0
	//h.Log.WithFields(logrus.Fields{
	//	"i":i,
	//	"a":a,
	//	"b":b,
	//}).Info()
	for ; i <= n; i++ {
		//h.Log.WithFields(logrus.Fields{
		//	"i":i,
		//	"a":a,
		//	"b":b,
		//}).Info()
		c = a + b
		a = b
		b = c
		h.Db.SetFibonacci(i,c)
	}

	return c
}