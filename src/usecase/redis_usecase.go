package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)


type RedisUsecase interface {
	AddKeyValueSet(key string, value string, expiration int) error
	GetValueByKey(key string) (string, error)
	DeleteValueByKey(key string) error
	ExistsByKey(key string) bool
}

type redisUsecase struct {
	RedisClient *redis.Client
}

func NewRedisUsecase(r *redis.Client) RedisUsecase{
	return &redisUsecase{r}
}

func (r2 *redisUsecase) AddKeyValueSet(key string, value string, expiration int) error{

	err := r2.RedisClient.Set(context.Background(), key, value, time.Duration(expiration)).Err()

	return err
}

func (r2 *redisUsecase) GetValueByKey(key string) (string,error) {
	val, err := r2.RedisClient.Get(context.Background(), key).Result()

	return val, err
}

func (r2 *redisUsecase) DeleteValueByKey(key string) error {

	return r2.RedisClient.Del(context.Background(), key).Err()
}

func (r2 *redisUsecase) ExistsByKey(key string) bool {

	err := r2.RedisClient.Exists(context.Background(), key).Err()

	if err != nil {
		return false
	}

	return true
}





