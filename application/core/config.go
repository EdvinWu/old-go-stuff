package core

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./home-task-tracker/configuration/")
	viper.SetConfigName("macos.config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type RedisConfig struct {
	Address    string
	Password   string
	DatabaseNo int
}

func GetDatabaseConfiguration() Database {
	var db Database

	err := viper.UnmarshalKey("database", &db)
	if err == nil {
		return db
	}
	panic(err)
}

func GetRedisConfiguration() RedisConfig {
	var redis RedisConfig

	err := viper.UnmarshalKey("redis", &redis)
	if err == nil {
		return redis
	}
	panic(err)
}

func GetCacheConfiguration(cacheName CacheName) map[string]string {
	return viper.GetStringMapString(string(cacheName))
}
