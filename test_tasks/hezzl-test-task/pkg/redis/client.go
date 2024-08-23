package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *Config) *redis.Client {
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
}
