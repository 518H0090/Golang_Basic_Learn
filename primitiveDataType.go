package main

import "fmt"

func main() {
	/*
		Boolean
		Number
		Text
	*/
	// Integer
	/*
		signed integer int8 -> int64
		unsigned integer uint8 -> uint64

		+ - * /
		& | ^ &^
		>> <<
	*/

	// Type Boolean
	var aBoo bool = false
	fmt.Printf("%v , %T", aBoo, aBoo)

	aBoo = 1 == 1

	fmt.Println(aBoo)

	// Type Number
	// Integer
	// (signed integer && usigned integer)
	// int8 int16 int32 int64 - uint8 uint16 uint32 uint64

	// Float
	/*
		float32
		float64
	*/
	var a float32 = 3.14
	fmt.Printf("%v %T", a, a)

	b := 13.7e42
	fmt.Println(b)
}
