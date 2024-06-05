package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func InitRedis(ctx context.Context) (*redis.Client, error) {
	config := ctx.Value("redis").(common.RedisConfiguration)

	cacheURI := fmt.Sprintf(
		"%s:%s",
		config.RedisHost,
		config.RedisPort,
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     cacheURI,
		Password: config.RedisPassword,
	})

	status := rdb.Ping(ctx)
	if status.Val() != "PONG" {
		return rdb, errors.New(status.String())
	}

	return rdb, nil
}
