package main

import "fmt"

func getUserInput() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func getInitialCurrency() string {
	var initialCurrency string
	fmt.Println("Выбери исходную валюту (USD / EUR / RUB):")
	for {
		initialCurrency = getUserInput()
		if initialCurrency == "USD" || initialCurrency == "EUR" || initialCurrency == "RUB" {
			break
		} else {
			fmt.Println("Некорректное значение. Допустимо: USD / EUR / RUB")
		}
	}
	return initialCurrency
}

func getSum() float64 {
	var userSum float64
	fmt.Println("Введите количество денег (от 1 до 1000000):")
	for {
		fmt.Scan(&userSum)
		if userSum > 1 && userSum <= 1000000 {
			break
		}
		fmt.Println("Неверная сумма. Допустимые значения (от 1 до 1000000):")
		continue

	}
	return userSum
}
func getTargetCurrency(initialCurrency string) string {
	var targetCurrency string
	var acceptedCurrencies string
	switch initialCurrency {
	case "USD":
		acceptedCurrencies = "EUR, RUB"
	case "RUB":
		acceptedCurrencies = "EUR, USD"
	case "EUR":
		acceptedCurrencies = "USD, RUB"
	}
	fmt.Printf("Выбери целевую валюту. Допустимые значения: %s.\n", acceptedCurrencies)
	for {
		targetCurrency = getUserInput()
		if initialCurrency == targetCurrency || (targetCurrency != "USD" && targetCurrency != "EUR" && targetCurrency != "RUB") {
			fmt.Printf("Введено недопустимое значение. Допустимые значения: %s.\n", acceptedCurrencies)
			continue
		} else {
			break
		}
	}
	return targetCurrency
}

func calculateConversion(sum float64, initialCurrency string, targetCurrency string) float64 {
	const USDtoEUR = 0.82
	const USDtoRUB = 63.52
	const EURtoUSD = 1 / USDtoEUR
	const RUBtoUSD = 1 / USDtoRUB
	const EURtoRUB = USDtoRUB / USDtoEUR
	const RUBtoEUR = 1 / EURtoRUB

	var result float64

	conversionKey := initialCurrency + "to" + targetCurrency

	switch conversionKey {
	case "USDtoEUR":
		result = sum * USDtoEUR
	case "USDtoRUB":
		result = sum * USDtoRUB
	case "EURtoUSD":
		result = sum * EURtoUSD
	case "EURtoRUB":
		result = sum * EURtoRUB
	case "RUBtoUSD":
		result = sum * RUBtoUSD
	case "RUBtoEUR":
		result = sum * RUBtoEUR
	}

	fmt.Printf("%.2f %s = %.2f %s\n", sum, initialCurrency, result, targetCurrency)
	return result
}

func main() {
	initialCurrency := getInitialCurrency()
	sum := getSum()
	targetCurrency := getTargetCurrency(initialCurrency)
	calculateConversion(sum, initialCurrency, targetCurrency)
}
