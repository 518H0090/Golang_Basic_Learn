package main

import (
	"fmt"
	"math"
)

/*
	Quy tắc đặt tên chung là camel - khi chung package
	nếu sử dụng trong nhiều package thì nên tường minh
	muốn package ngoài sử dụng dc thì để viết hoa ví dụ NumberA
*/

const cConst = 10

func main() {

	// Constant - hằng số - Không thể thay đổi
	// Phải có sẵn giá trị khi khởi tạo
	const aConst string = "Hello World"
	fmt.Printf("%v %T \n", aConst, aConst)

	const bConst = 2
	fmt.Printf("%v %T \n", bConst, bConst)

	fmt.Printf("%v %T \n", cConst, cConst)

	var aNumber float32 = float32(math.Sqrt(9))
	fmt.Printf("%v %T \n", aNumber, aNumber)

	// Enum - liệt kê
	// Tập hợp các hằng số

	const (
		red    = 0
		yellow = 1
		blue   = 2
	)

	fmt.Printf("%v %T \n", red, red)
	fmt.Printf("%v %T \n", yellow, yellow)
	fmt.Printf("%v %T \n", blue, blue)

	fmt.Println()

	// Tự động tạo giá trị cho Enum bằng iota
	const (
		haha = iota
		hoho
		hihi
		hehe
	)

	fmt.Printf("%v %T \n", haha, haha)
	fmt.Printf("%v %T \n", hoho, hoho)
	fmt.Printf("%v %T \n", hihi, hihi)
	fmt.Printf("%v %T \n", hehe, hehe)

	fmt.Println()

	const (
		defaultE = iota
		E1
		E2
		E3
	)

	fmt.Printf("%v %T \n", E1, E1)
	fmt.Printf("%v %T \n", E2, E2)
	fmt.Printf("%v %T \n", E3, E3)

	// _  - biến không lưu vào đâu cả
	const (
		_ = iota
		D1
		D2
		D3
	)

	fmt.Printf("%v %T \n", D1, D1)
	fmt.Printf("%v %T \n", D2, D2)
	fmt.Printf("%v %T \n", D3, D3)

	const (
		_ = 1<<iota + 5
		F1
		F2
		F3
	)

	fmt.Printf("%v %T \n", F1, F1)
	fmt.Printf("%v %T \n", F2, F2)
	fmt.Printf("%v %T \n", F3, F3)

}
