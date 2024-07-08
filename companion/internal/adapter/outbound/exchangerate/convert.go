package exchangerate

import (
	"companion/internal/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const AppId = "584733aac8394c2597e7f5c673014c62"

var baseCurrency = "eur"
var rates = map[string]float64{}

func New() {

	c := config.Get()
	if c.Settings.BaseCurrency == "" {
		c.Settings.BaseCurrency = "eur"
		config.Set(c)
	}
	baseCurrency = c.Settings.BaseCurrency

	updateRates()

	config.Subscribe(func(c config.Config) {
		if c.Settings.BaseCurrency == "" {
			c.Settings.BaseCurrency = "eur"
			config.Set(c)
			return
		}

		if baseCurrency != c.Settings.BaseCurrency {
			baseCurrency = c.Settings.BaseCurrency
			updateRates()
		}
	})
}

func Convert(fromAmount float64, fromCurrency string) float64 {
	fromCurrency = strings.ToLower(fromCurrency)
	baseCurrency = strings.ToLower(baseCurrency)

	return rates[fromCurrency] * fromAmount

}

func updateRates() {
	log.Println("Downloading exchange rates for", baseCurrency)
	resp, err := http.Get(fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%v", AppId))
	if err != nil {
		log.Println("Error downloading exchange rates", err)
		return
	}
	data := dto{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error decoding exchange rates", err)
		return
	}

	upperCaseBaseCurrency := strings.ToUpper(baseCurrency)
	conversions := map[string]float64{}

	toCurrencyConversion, exists := data.Rates[upperCaseBaseCurrency]
	if !exists {
		log.Println("Cant handle unit", upperCaseBaseCurrency)
		return
	}

	for currency, factor := range data.Rates {
		//log.Println("Converting", currency, "to", upperCaseBaseCurrency, "with rate", factor, "and base rate", toCurrencyConversion)
		conversion := toCurrencyConversion / factor

		conversions[strings.ToLower(currency)] = conversion
	}

	rates = conversions
}

type dto struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
