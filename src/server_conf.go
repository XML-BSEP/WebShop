package main

import (
	"github.com/spf13/viper"
	"log"
)

func init_viper() {
	viper.SetConfigFile(`src/configurations/serverconfig.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func NewSSLServer() (string, string, string){
	init_viper()
	address := viper.GetString(`server.address`)
	port := viper.GetString(`server.port`)

	certificate := viper.GetString(`certificate`)
	certificate_key := viper.GetString(`certificate_key`)

	return address + ":" + port, certificate, certificate_key
}

func NewHttpServer() string {
	init_viper()
	address := viper.GetString(`httpserver.address`)
	port := viper.GetString(`httpserver.port`)


	return address + ":" + port
}