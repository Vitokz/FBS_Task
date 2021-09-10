package handler_test

import (
	"github.com/Vitokz/Task/handler"
	"github.com/Vitokz/Task/repository"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRest_Fibbonaci(t *testing.T) {
	log := logrus.Logger{}
	hndlr := handler.Handler{
		Log: &log,
		Db: repository.DbConn(),
	}

	resp,err:=hndlr.Fibonacci(1,10)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equalf(t, "[1] = 0, [2] = 1, [3] = 1, [4] = 2, [5] = 3, [6] = 5, [7] = 8, [8] = 13, [9] = 21, [10] = 34",resp.Numbers,"failed")
}
