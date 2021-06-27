package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"web-shop/domain"
)

const (
	cart = "cart"
)
type RedisUsecase interface {
	AddKeyValueSet(key string, value string, expiration int) error
	GetValueByKey(key string) (string, error)
	DeleteValueByKey(key string) error
	ExistsByKey(key string) bool
	SetToken(key string, value string, time time.Duration) error
	CheckUsername(username string) bool
}

type redisUsecase struct {
	RedisClient *redis.Client
	RegisteredUserRepository domain.RegisteredShopUserRepository
}

func NewRedisUsecase(r *redis.Client, usecase domain.RegisteredShopUserRepository) RedisUsecase{
	return &redisUsecase{r, usecase}
}


func (r2 *redisUsecase) SetToken(key string, value string, time time.Duration) error {
	err := r2.RedisClient.Set(context.Background(), key, value, time).Err()
	return err
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
	err := r2.RedisClient.Get(context.Background(), key).Err()
	if err == redis.Nil{
		return false
	}
	return true
}


func (r2 *redisUsecase) CheckUsername(username string) bool {
	var userObj domain.UserRegistrationRequest
	var cursor uint64
	keys, _, err := r2.RedisClient.Scan(context.Background(), cursor, "*", 10).Result()
	if err != nil {
		return false
	}

	for _, key := range keys {
		value, err := r2.GetValueByKey(key)
		if err == nil {
			err = json.Unmarshal([]byte(value), &userObj)

			if err == nil {
				if userObj.Username == username {
					return true
				}
			}
		}
	}

	fmt.Println(keys)
	return false
}





