package main

import (
	"fmt"
)

type Country string

const (
	GERMANY Country = "Germany"
	FRANCE  Country = "France"
	TURKEY  Country = "Turkey"
	ENGLAND Country = "England"
)

func printCountry(country Country) {
	fmt.Println(country)
}

func main() {
	printCountry(TURKEY)
	printCountry(GERMANY)
}
