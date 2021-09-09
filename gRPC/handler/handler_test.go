package handler_test

import (
"context"
"github.com/Vitokz/Task/gRPC/config"
"github.com/Vitokz/Task/gRPC/handler"
"github.com/Vitokz/Task/gRPC/repository"
"github.com/stretchr/testify/assert"
"testing"
)

func TestRest_Fibbonaci(t *testing.T) {
	cfg, err := config.Parse("/home/vitoo/go/src/Task/gRPC/config/config.toml")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	hndlr := handler.Handler{
		Config: cfg,
		Db: repository.DbConn(cfg.DateBase.Port),
	}

	resp,err:=hndlr.Fibonacci("0","10",context.Background())
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	_=resp
	assert.Equalf(t, "0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55",resp.Numbers,"failed")
}