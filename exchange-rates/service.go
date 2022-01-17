package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CurrencyRate struct {
	FromCurrencyName string  `json:"from-currency"`
	ToCurrencyName   string  `json:"to-currency"`
	Rate             float64 `json:"rate"`
}

type ExchangeRateService struct {
	ExchangeRepo *ExchangeRepo
}

func (c *ExchangeRateService) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func (c *ExchangeRateService) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	from := string(vars["from"])
	to := string(vars["to"])
	rate := &CurrencyRate{
		FromCurrencyName: from,
		ToCurrencyName:   to,
		Rate:             c.ExchangeRepo.FetchExchangeRate(from, to)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)

}
