package main

import "fmt"

func main() {
	// MAIN TYPES:
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Corey"
	var age = 35
	const isCool = true

	// shorthad var declartion (can't declare like this outside of function)
	// name := "Corey"
	size := 1.3
	
	name, email := "Corey", "corey@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
