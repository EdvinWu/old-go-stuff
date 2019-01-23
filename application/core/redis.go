package core

import (
	"github.com/go-redis/redis"
	"log"
	"time"
	"fmt"
	"strconv"
)

const (
	DAY                    = time.Hour * 24
	PARENT_CACHE CacheName = "parent_cache"
)

type CacheName string

type TimeUnit struct {
	value string
	unit  string
}

func init() {
	GetRedisConnection()
}
func WriteToRedis(key string, value interface{}, name CacheName) error {
	client := GetRedisConnection()
	duration, err := getDuration(GetCacheConfiguration(name))
	if err == nil {
		set := client.Set(key, value, duration)
		fmt.Printf("%#v\n %T", set, set)
	}
	return err
}

func GetRedisConnection() *redis.Client {
	config := GetRedisConfiguration()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DatabaseNo,
	})
	_, err := client.Ping().Result()
	if err == nil {
		log.Println("Successfully connected to redis")
		return client
	}
	panic(err)
}

func getDuration(cacheMap map[string]string) (dur time.Duration, err error) {
	unit := cacheMap["unit"]
	switch s := unit; s {
	case "SECOND":
		dur = time.Second
		break
	case "MINUTE":
		dur = time.Minute
		break
	case "DAY":
		dur = DAY
		break
	default:
		dur = time.Hour
	}
	numVal, err := strconv.Atoi(cacheMap["value"])
	if err == nil {
		return dur * time.Duration(numVal), nil
	}
	return dur, err
}
