package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	log.Println("Get Redis Instance")
	redisClient := GetRedis()
	log.Println("Got Redis Instance")

	exchangeRepo := &ExchangeRepo{redisClient: redisClient}
	currencyService := &ExchangeRateService{ExchangeRepo: exchangeRepo}
	r.HandleFunc("/exchange-rates/health", currencyService.HealthCheck)
	r.HandleFunc("/exchange-rates/{from}/to/{to}", currencyService.GetExchangeRate)
	http.Handle("/", r)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr:         "0.0.0.0:9001",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
