package auth

import (
	"fmt"
	"strconv"
	"time"
	"web-shop/domain/auth"
	"web-shop/usecase"
)

type AuthInterface interface {
	CreateAuth(uint, *auth.TokenDetails) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*auth.AccessDetails) error
	AlreadyDeleted(metadata *auth.AccessDetails) bool
}

type RedisUseCase struct {
	redisUseCase usecase.RedisUsecase
}

func (tk *RedisUseCase) AlreadyDeleted(metadata *auth.AccessDetails) bool {
	if tk.redisUseCase.ExistsByKey(metadata.TokenUuid) {
		return true
	}
	return false
}

var _ AuthInterface = &RedisUseCase{}

func NewAuth(redisUseCase usecase.RedisUsecase) *RedisUseCase {
	return &RedisUseCase{redisUseCase : redisUseCase}
}


func (tk *RedisUseCase) CreateAuth(userid uint, td *auth.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err := tk.redisUseCase.SetToken(td.TokenUuid, strconv.Itoa(int(userid)), at.Sub(now))

	if tk.redisUseCase.ExistsByKey(td.TokenUuid) {
		return nil
	}
	if err != nil {
		return err
	}

	err = tk.redisUseCase.SetToken(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now))
	if err != nil {
		return err
	}

	return nil
}


func (tk *RedisUseCase) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := tk.redisUseCase.GetValueByKey(tokenUuid)
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}


func (tk *RedisUseCase) DeleteTokens(authD *auth.AccessDetails) error {

	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)

	err := tk.redisUseCase.DeleteValueByKey(authD.TokenUuid)
	if err != nil {
		return err
	}

	err = tk.redisUseCase.DeleteValueByKey(refreshUuid)
	if err != nil {
		return err
	}

	return nil
}

func (tk *RedisUseCase) DeleteRefresh(refreshUuid string) error {
	err := tk.redisUseCase.DeleteValueByKey(refreshUuid)
	if err != nil {
		return err
	}
	return nil
}