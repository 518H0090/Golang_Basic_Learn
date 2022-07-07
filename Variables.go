package main

import (
	"fmt"
	"strconv"
)

// Global Scope
var (
	n int = 10
	m int = 20
)

var (
	stI string = "HAHA"
)

func main() {
	// Block Scope
	// Khai báo tường minh
	// Khai báo biến cách 1
	var hi int
	hi = 1
	fmt.Println("Hello World")
	fmt.Print(hi)

	// Cách 2
	var hi2 int = 10
	fmt.Println(hi2)

	// Cách 3
	// Khai báo ngầm định
	hi3 := 20
	fmt.Println(hi3)

	// Xuất ra màn hình
	// %v là giá trị , %T là kiểu dữ liệu
	fmt.Printf("%v , %T", hi3, hi3)

	fmt.Println(n)
	fmt.Println(m)

	// Convert Type

	var numValue1 float32 = 10.5
	fmt.Println(numValue1)

	var numValue2 int = int(numValue1)

	fmt.Println(numValue2)

	var iNumber int = 65

	var characterNumber string = string(iNumber)
	fmt.Println(characterNumber)

	var sNumber string = strconv.Itoa(iNumber)

	fmt.Println(sNumber)

}
