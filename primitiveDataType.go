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
	fmt.Printf("%v %T \n", a, a)

	b := 13.7e42
	fmt.Println(b)

	// Complex complex64 complex128 - 1 phần thực 1 phần ảo
	var aComplex complex64 = 1 + 2i
	fmt.Println(aComplex)
	fmt.Printf("%v %T \n", real(aComplex), real(aComplex))
	fmt.Printf("%v %T \n", imag(aComplex), imag(aComplex))

	var bComplex complex128 = complex(5, 12)
	fmt.Printf("%v %T \n", bComplex, bComplex)

	// Text
	/*
		String
		Byte
		Rune
	*/

	// String
	var aString string = "Hello World"
	fmt.Printf("%v %T  \n", aString, aString)

	var bString string = "Complex 123"

	fmt.Printf("%v %T  \n", aString+bString, aString+bString)

	// Byte -> UTF-8 -> dùng lưu trữ alphabet
	// Chia chuoi a thanh nhieu byte nhỏ, 8bit -> 1byte
	var aByteArray []byte = []byte(aString)

	fmt.Printf("%v %T  \n", aByteArray, aByteArray)

	var bByte byte = 'H'
	fmt.Printf("%v %T  \n", bByte, bByte)

	fmt.Printf("%v %T  \n", string(bByte), bByte)

	// Rune -> UTF-32 -> dùng lưu trữ các dạng chữ phức tạp hơn
	var aRune rune = 'e'
	fmt.Printf("%v %T  \n", aRune, aRune)

	fmt.Printf("%v %T  \n", string(aRune), aRune)
}
