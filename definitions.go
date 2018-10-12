package main

import "fmt"

var classVariable = "Class Variable"

const constVariable = "Const Variable"
const pi = 3.14

var (
	name    = "Burak"
	surname = "İren"
)

func main() {

	// değişken ismi küçük harfle başlarsa private, büyük harfle başlarsa public
	var message string
	message = "Merhaba Go!"
	fmt.Println(message)

	var message2 string = "Test 3"
	var message3 = "Test 4"
	println(message2)
	println(message3)

	var a, b, c int = 3, 4, 5
	println(a)
	println(b)
	println(c)

	// Tek satırda farklı tiplere değişken tanımlama
	var d, e, f = 3, true, 4.5
	println(d, e, f)

	message4 := "kısa değişken tanımlama" // kısa değişken tanımalama, sadece function body içinde tanımlanır, class değişkeni olursa hata verir
	println(message4)

	message5, message6 := "message5", true
	println(message5, message6)

	message7 := "B"
	message8 := 'B' // Tek tırnak olunca char değerini yazıyor
	message9 := `B`
	println(message7, message8, message9)

	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	println(myFloat32)
	println(myComplex)

	println(classVariable)
	println(constVariable)
	println(pi)

	println(name, surname)

}
