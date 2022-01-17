package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

type ExchangeRepo struct {
	redisClient *redis.Client
}

func GetRedis() *redis.Client {
	// Create Redis Client
	var (
		host     = os.Getenv("REDIS_HOST")
		port     = os.Getenv("REDIS_PORT")
		password = os.Getenv("REDIS_PASSWORD")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	// Check we have correctly connected to our Redis server
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (e *ExchangeRepo) FetchExchangeRate(from string, to string) float64 {
	key := from + "_TO_" + to
	rate, _ := e.redisClient.Get(key).Float64()
	return rate

}
