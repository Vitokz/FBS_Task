package repository

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type Db struct {
	Cache *cache.Cache
}
const port = "6379"
var ctx=context.Background()

func (d *Db) SetValue(key string, value string) error {
	if err := d.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	}); err != nil {
		return err
	}
	return nil
}

func (d *Db) GetValue(key string) (string, error) {
	var value string
	if err := d.Cache.Get(ctx, key, &value); err != nil {
		return value, err
	}
	return value, nil
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
		Cache: mycache,
	}

	db.SetValue("0","0")
	db.SetValue("1","1")

	return db
}
