package exchangerate

import (
	"companion"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const AppId = "584733aac8394c2597e7f5c673014c62"

func New(targetCurrency string) (companion.CurrencyConverter, error) {

	rates, err := getRates(targetCurrency)

	if err != nil {
		return nil, err
	}

	return &converter{
		rates: rates,
	}, nil
}

type converter struct {
	rates map[string]float64
}

func (c *converter) Convert(fromAmount int, fromCurrency string, toCurrency string) int {
	fromCurrency = strings.ToLower(fromCurrency)
	toCurrency = strings.ToLower(toCurrency)

	return int(c.rates[toCurrency] * float64(fromAmount))

}

func getRates(toCurrency string) (map[string]float64, error) {

	log.Println("Downloading exchange rates for", toCurrency)
	resp, err := http.Get(fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%v", AppId))
	if err != nil {
		return nil, err
	}
	data := ApiResponse{}
	json.NewDecoder(resp.Body).Decode(&data)

	toCurrency = strings.ToUpper(toCurrency)
	conversions := map[string]float64{}

	toCurrencyConversion, exists := data.Rates[toCurrency]
	if !exists {
		return nil, fmt.Errorf("cant handle unit %v", toCurrency)
	}

	for currency, factor := range data.Rates {
		conversion := toCurrencyConversion * factor

		conversions[strings.ToLower(currency)] = conversion
	}

	return conversions, nil
}

type ApiResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
