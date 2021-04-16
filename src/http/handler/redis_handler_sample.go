package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"web-shop/infrastructure/dto"
	"web-shop/usecase"
)

type RedisHandlerSample interface {
	AddKeyValueSet(c echo.Context) error
	GetValueByKey(c echo.Context) error
}

type redisHandlerSample struct {
	RedisUsecase usecase.RedisUsecase
}


func NewRedisHandlerSample(r usecase.RedisUsecase) RedisHandlerSample{
	return &redisHandlerSample{r}
}

func (r *redisHandlerSample) AddKeyValueSet(c echo.Context) error {

	decoder := json.NewDecoder(c.Request().Body)

	var newUser dto.NewUser
	
	err := decoder.Decode(&newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}

	expiration  := 1000000000*3600*2
	err2 := r.RedisUsecase.AddKeyValueSet(newUser.Username, newUser.Password, expiration)

	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, "OK")
}


func (r *redisHandlerSample) GetValueByKey(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)

	var key dto.RedisKey

	err := decoder.Decode(&key)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error")
	}

	val, err := r.RedisUsecase.GetValueByKey(key.Key)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error")
	}
	return c.JSON(http.StatusOK, val)
}

