package main

import "fmt"

func main() {
	// Integer
	/*
		signed integer int8 -> int64
		unsigned integer uint8 -> uint64

		+ - * /
		& | ^ &^
		>> <<
	*/

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
