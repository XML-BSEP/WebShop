package database

import (
	"web-shop/infrastructure/security/auth"
	"github.com/go-redis/redis/v7"
)


type RedisService struct {
	Auth   auth.AuthInterface
	Client *redis.Client
}

func NewRedisDB(host, port, password string) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &RedisService{
		Auth:   auth.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
