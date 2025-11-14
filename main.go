package main

import "fmt"

func main() {
	const USDtoEUR = 0.82
	const USDtoRUB = 63.52
	var EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
