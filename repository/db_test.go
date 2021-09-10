package repository_test

import (
	"github.com/Vitokz/Task/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbConn(t *testing.T) {
	key := "-100"
	val := "-100"
	conn := repository.DbConn()
	err := conn.SetValue("-100", "-100")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	value, err := conn.GetValue(key)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, val,value,"error")
}
