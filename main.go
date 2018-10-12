package main

import "fmt"

func main() {

	// değişken ismi küçük harfle başlarsa private, büyük harfle başlarsa public
	var message string
	message = "Merhaba Go!"
	fmt.Println(message)

	var message2 string = "Test 3"
	var message3 = "Test 4"
	fmt.Println(message2)
	fmt.Println(message3)

	var a, b, c int = 3, 4, 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d, e, f = 3, true, 4.5
	fmt.Println(d, e, f)

}
