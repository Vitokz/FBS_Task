package repository_test

import (
	"context"
	"github.com/Vitokz/Task/Rest/config"
	"github.com/Vitokz/Task/Rest/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbConn(t *testing.T) {
	cfg, err := config.Parse("/home/vitoo/go/src/Task/Rest/config/config.toml")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	key := "-100"
	val := "-100"
	conn := repository.DbConn(cfg.DateBase.Port)
	err = conn.SetValue(context.Background(), "-100", "-100")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	value, err := conn.GetValue(context.Background(), key)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, val,value,"error")
}
