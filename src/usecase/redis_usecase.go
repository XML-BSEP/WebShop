package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)


type RedisUsecase interface {
	AddKeyValueSet(key string, value string, expiration int) error
	GetValueByKey(key string) (string, error)
}

type redisUsecase struct {
	RedisClient *redis.Client
}

func NewRedisUsecase(r *redis.Client) RedisUsecase{
	return redisUsecase{r}
}

func (r2 redisUsecase) AddKeyValueSet(key string, value string, expiration int) error{

	err := r2.RedisClient.Set(context.Background(), key, value, time.Duration(expiration)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r2 redisUsecase) GetValueByKey(key string) (string,error) {
	val, err := r2.RedisClient.Get(context.Background(), key).Result()

	if err != nil {
		return "", err
	}
	return val, nil
}






