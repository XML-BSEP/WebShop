package redisdb

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func init_viper() {
	viper.SetConfigFile(`src/configurations/dbconfig.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}


}

func NewReddisConn() *redis.Client {
	init_viper()
	//address := viper.GetString(`server.address`)
	//port := viper.GetString(`server.port`)

	//redisAddr := address + ":" + port
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})


}
