package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	redisClient := GetRedis()
	exchangeRepo := &ExchangeRepo{redisClient: redisClient}
	currencyService := &CurrencyService{ExchangeRepo: exchangeRepo}
	r.HandleFunc("/currency-service/health", currencyService.HealthCheck)
	r.HandleFunc("/currency-service/{from}/to/{to}", currencyService.GetExchangeRate)
	http.Handle("/", r)
	log.Println("Starting service...")
	http.ListenAndServe(":8080", r)

}
