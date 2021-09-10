package handler_test

import (
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRest_Fibbonaci(t *testing.T) {
	hndlr := handler.Handler{
		Db: repository.DbConn(),
	}

	resp,err:=hndlr.Fibonacci(0,10)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equalf(t, "[0] = 0, [1] = 1, [2] = 1, [3] = 2, [4] = 3, [5] = 5, [6] = 8, [7] = 13, [8] = 21, [9] = 34, [10] = 55",resp.Numbers,"failed")
}
