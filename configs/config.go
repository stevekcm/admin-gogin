package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnv()
}

type envConfigs struct {
	MongoUri  string `mapstructure:"MONGO_URI"`
	MongoName string `mapstructure:"MONGO_NAME"`
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPass string `mapstructure:"REDIS_PASS"`
}

func loadEnv() (config *envConfigs) {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Failed to load .env file", err.Error())
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Println("Failed to unmarshal .env file")
		panic(err)
	}

	return
}
