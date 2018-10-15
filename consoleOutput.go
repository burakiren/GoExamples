package main

import (
	"fmt"
)

func main() {

	str1 := "Bu bir deneme yazisi"
	str2 := "Bu bir test yazisi"
	str3 := "Bu bir harakiri yazisi"
	aNumber := 54
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)
	fmt.Println(stringLength)
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	fmt.Printf("Data Types : %T , %T, %T , %T", str1, str2, aNumber, isTrue)
	fmt.Println("")

}
