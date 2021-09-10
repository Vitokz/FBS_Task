package repository_test

import (
	"github.com/Vitokz/Task/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbConn(t *testing.T) {
	val := "-100"
	conn := repository.DbConn()
	err := conn.SetFibonacci(-100, -100)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	value, err := conn.GetFibonacci(-100)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, val,value,"error")
}
