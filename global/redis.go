package global

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"sync"
)

var (
	rdb *redis.Client

	rdbOnce sync.Once
)

func RDB() *redis.Client {
	return rdb
}

func CloseRDB() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}

func initRedis() {
	rdbOnce.Do(func() {
		_db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			panic(err.Error())
		}

		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       _db,
		})
	})
}
