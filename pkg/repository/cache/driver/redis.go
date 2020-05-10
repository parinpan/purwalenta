package driver

import (
	"github.com/go-redis/redis"
	"github.com/parinpan/purwalenta/pkg/config"
)

func GetRedisDriver() (*redis.Client, error) {
	cfg := config.GetConfig().Database.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       0,
	})

	if _, err := client.Ping().Result(); nil != err {
		return client, err
	}

	return client, nil
}
