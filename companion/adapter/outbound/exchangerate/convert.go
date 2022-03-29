package exchangerate

import (
	"companion"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func New(targetCurrency string) (companion.CurrencyConverter, error) {
	rates, err := getRates(targetCurrency)
	if err != nil {
		return nil, err
	}
	return &converter{
		rates: *rates,
	}, nil
}

type converter struct {
	rates ConversionRates
}

func (c *converter) Convert(fromAmount int, fromCurrency string, toCurrency string) int {

	fromCurrency = strings.ToLower(fromCurrency)

	return int(c.rates[fromCurrency].InverseRate * float64(fromAmount))

}

func getRates(toCurrency string) (*ConversionRates, error) {

	log.Println("Downloading exchange rates for", toCurrency)
	resp, err := http.Get(fmt.Sprintf("http://www.floatrates.com/daily/%v.json", strings.ToLower(toCurrency)))
	if err != nil {
		return nil, err
	}
	rates := &ConversionRates{}
	json.NewDecoder(resp.Body).Decode(rates)

	return rates, nil
}

type ConversionRates map[string]struct {
	Code        string  `json:"code"`
	AlphaCode   string  `json:"alphaCode"`
	NumericCode string  `json:"numericCode"`
	Name        string  `json:"name"`
	Rate        float64 `json:"rate"`
	Date        string  `json:"date"`
	InverseRate float64 `json:"inverseRate"`
}
