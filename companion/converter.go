package companion

type CurrencyConverter interface {
	Convert(fromAmount int, fromCurrency string, toCurrency string) int
}
