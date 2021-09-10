package repository

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Db struct {
	cache *cache.Cache
}
const port = "6379"
var ctx=context.Background()

func (d *Db) setValue(key string, value string) error {
	if err := d.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		return err
	}
	return nil
}

func (d *Db) getValue(key string) (string, error) {
	var value string
	if err := d.cache.Get(ctx, key, &value); err != nil {
		return value, err
	}
	return value, nil
}

func (d *Db) GetFibonacci(key int) (string,error) {
	return d.getValue(strconv.Itoa(key))
}

func (d *Db) SetFibonacci(key int,value int) error {
	return d.setValue(strconv.Itoa(key),strconv.Itoa(value))
}
func DbConn() *Db {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":" + port,
		},
		DB: 1,
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	db:= &Db{
		cache: mycache,
	}

	db.SetFibonacci(1,0)
	db.SetFibonacci(2,1)

	return db
}
