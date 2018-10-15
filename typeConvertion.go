package main

import (
	"strconv"
)

func main() {

	var myString = "54"
	number, _ := strconv.Atoi(myString) // String to int
	var str = strconv.Itoa(number)      // int to String
	println(str)

	var i int = 55
	var f float64 = float64(i) // int to float
	var u uint = uint(f)       // float to uint

	println(u)

}
