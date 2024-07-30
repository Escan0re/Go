package main

import (
	"fmt"
)

func main() {

	sourceCurrency, amountOfFunds, targetCurrency := getUserInput()
	conversationResult := currencyConversion(sourceCurrency, amountOfFunds, targetCurrency)
	fmt.Printf("Исходная валюта: %s, Сумма средств в исходной валюте: %.0f, Таргетная валюта: %s\n", sourceCurrency, amountOfFunds, targetCurrency)
	fmt.Printf("Результат конвертации: %.2f\n", conversationResult)

}

func getUserInput() (string, float64, string) {
	var sourceCurrency string
	var amountOfFunds float64
	var targetCurrency string

	for {
		fmt.Print("Введите исходную валюту, например EUR, USD, RUB: ")
		fmt.Scanln(&sourceCurrency)
		if sourceCurrency == "" || (sourceCurrency != "rub" && sourceCurrency != "eur" && sourceCurrency != "usd") {
			continue
		}
		break
	}
	for {
		fmt.Print("Введите сумму средств в исходной валюте: ")
		fmt.Scanln(&amountOfFunds)
		if amountOfFunds <= 0 {
			fmt.Println("Ошибка. Нельзя вводить отрицательные значения!")
			continue
		}
		break
	}
	for {
		fmt.Print("Введите таргетную валюту, например EUR, USD, RUB: ")
		fmt.Scanln(&targetCurrency)
		if targetCurrency == "" || (targetCurrency != "rub" && targetCurrency != "eur" && targetCurrency != "usd") {
			continue
		}
		if targetCurrency == sourceCurrency {
			fmt.Print("Исходная и таргетная валюты совпадают, измени таргетную валюту: ")
			continue
		}
		break
	}
	return sourceCurrency, amountOfFunds, targetCurrency
}

func currencyConversion(sourceCurrency string, amountOfFunds float64, targetCurrency string) float64 {
	const (
		eurToUsd = 1.08
		usdToEur = 0.93
		rubToEur = 0.013
		eurToRub = 76.95
		usdToRub = 71.43
		rubToUsd = 0.014
	)
	var result float64
	switch {
	case sourceCurrency == "eur" && targetCurrency == "usd":
		result = amountOfFunds * eurToUsd
	case sourceCurrency == "usd" && targetCurrency == "eur":
		result = amountOfFunds * usdToEur
	case sourceCurrency == "rub" && targetCurrency == "eur":
		result = amountOfFunds * rubToEur
	case sourceCurrency == "eur" && targetCurrency == "rub":
		result = amountOfFunds * eurToRub
	case sourceCurrency == "usd" && targetCurrency == "rub":
		result = amountOfFunds * usdToRub
	case sourceCurrency == "rub" && targetCurrency == "usd":
		result = amountOfFunds * rubToUsd
	}
	return result
}
