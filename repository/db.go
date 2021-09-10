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

func (d *Db) SetValue(ctx context.Context, key string, value string) error {
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

func (d *Db) GetValue(ctx context.Context, key string) (string, error) {
	var value string
	if err := d.Cache.Get(ctx, key, &value); err != nil {
		return value, err
	}
	return value, nil
}

func DbConn(port string) *Db {
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

	return &Db{
		Cache: mycache,
	}
}
