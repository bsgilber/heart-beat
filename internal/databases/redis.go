package databases

import "github.com/go-redis/redis/v8"

func ConnectClient() *redis.Client {
	opt, err := redis.ParseURL("redis://redis:6379")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	return rdb
}
