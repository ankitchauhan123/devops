package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type ExchangeRateResp struct {
	FromCurrencyName string  `json:"from-currency"`
	ToCurrencyName   string  `json:"to-currency"`
	Rate             float64 `json:"rate"`
	Amount           float64 `json:"amount"`
}

type CurrencyRate struct {
	FromCurrencyName string  `json:"from-currency"`
	ToCurrencyName   string  `json:"to-currency"`
	Rate             float64 `json:"rate"`
}

type CurrencyService struct {
	httpClient *http.Client
}

func (c *CurrencyService) convert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	from := string(vars["from"])
	to := string(vars["to"])
	log.Println("Amount:", vars["amount"])
	amount, _ := strconv.ParseFloat(string(vars["amount"]), 64)

	exchangeRate := c.FetchExchangeRate(from, to)
	rate := &ExchangeRateResp{
		FromCurrencyName: from,
		ToCurrencyName:   to,
		Rate:             exchangeRate.Rate,
		Amount:           amount * exchangeRate.Rate,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)

}

func (c *CurrencyService) FetchExchangeRate(from string, to string) *CurrencyRate {
	host := os.Getenv("SERVICE_HOST")
	port := os.Getenv("SERVICE_PORT")

	url := "http://" + host + ":" + port + "/exchange-rates/" + from + "/to/" + to
	log.Println("Fetching exchange rate from url", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	if err != nil {
		panic(err.Error())
	}

	cr := &CurrencyRate{}
	json.Unmarshal(body, cr)
	log.Println("Data Fetched:", cr, ":", cr.Rate)

	return cr
}
