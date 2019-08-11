package config

import (
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	Rabbit *rabbit
)

type rabbit struct {
	User     string
	Password string
	Host     string
	Port     string
}

func init() {
	config = viper.New()
	config.SetConfigFile("./config/config.yml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	initRabbitConfig()
}

func initRabbitConfig() {
	conf := config.GetStringMapString("RabbitMQ")
	Rabbit = &rabbit{
		User:     conf["user"],
		Password: conf["password"],
		Host:     conf["host"],
		Port:     conf["port"],
	}
}
