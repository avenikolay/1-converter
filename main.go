package main

import "fmt"

func getUserInput() string {
	var userInput string
	fmt.Scan(userInput)
	return userInput
}

func calculateConversion(sum float64, initialCurrency string, targetCurrency string) {
}

func main() {
	const USDtoEUR = 0.82
	const USDtoRUB = 63.52
	var EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
