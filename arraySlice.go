package main

import (
	"fmt"
)

func main() {
	// Basic
	// Kích thước của mảng luôn cố định kích thước khi khởi tạo, đôi khi sử dụng không hết => lãng phí bộ nhớ , muốn mở rộng không được => hạn chế tùy chỉnh.
	pointStudent1 := 7
	pointStudent2 := 8
	pointStudent3 := 10

	fmt.Printf("%v %T\n", pointStudent1, pointStudent1)
	fmt.Printf("%v %T\n", pointStudent2, pointStudent2)
	fmt.Printf("%v %T\n", pointStudent3, pointStudent3)

	// Array - tập hợp các biến có cùng kiểu dữ liệu
	points := [3]int{7, 5, 10}
	fmt.Printf("%v %T\n", points, points)

	var points2 []int = []int{7, 5, 20, 20}
	fmt.Printf("%v %T\n", points2, points2)

	points3 := [...]int{2, 03, 503, 25}
	fmt.Printf("%v %T\n", points3, points3)

	var points4 [3]int
	points4[0] = 2
	points4[1] = 5
	points4[2] = 10
	fmt.Printf("%v %T\n", points4, points4)
	fmt.Printf("%v %T\n", points4[2], points4[2])

	// Length Array
	fmt.Println(len(points4))

	// Copy Array
	a := [3]int{2, 5, 10}
	b := a

	b[1] = 50

	fmt.Println(a)
	fmt.Println(b)

	// Copy Array use Pointer
	aPoint := [3]int{2, 5, 10}
	bPoint := &aPoint

	bPoint[1] = 50

	fmt.Println(aPoint)
	fmt.Println(*bPoint)

	// Slice
	// Slice mang đặc điểm y như Array nhưng nó có kích thước Dynamic nên tùy chỉnh được
	// Mảng tập hợp con trỏ
	aSlice := []int{2, 5, 6, 10, 20, 40}
	fmt.Println(aSlice)

	bSlice := aSlice
	bSlice[3] = 75
	fmt.Println(bSlice)

	// Len Slice
	fmt.Println(len(bSlice))

	// cap - sức chứa
	fmt.Println(cap(aSlice))

	// Thay đổi chiều dài của Slice
	// Cắt chuỗi
	aEx := []int{2, 5, 10, 12, 45, 50}
	bEx := aEx[:]
	cEx := aEx[3:]
	dEx := aEx[:5]
	eEx := aEx[3:5]

	eEx[1] = 100

	fmt.Printf("%v %v %v\n", aEx, len(aEx), cap(aEx))
	fmt.Printf("%v %v %v\n", bEx, len(bEx), cap(bEx))
	fmt.Printf("%v %v %v\n", cEx, len(cEx), cap(cEx))
	fmt.Printf("%v %v %v\n", dEx, len(dEx), cap(dEx))
	fmt.Printf("%v %v %v\n", eEx, len(eEx), cap(eEx))

	//make - dùng make để khai báo slice xác định được len và cap
	aSl := make([]int, 10)
	fmt.Printf("%v %v %v\n", aSl, len(aSl), cap(aSl))

	bSl := make([]int, 10, 20)
	fmt.Printf("%v %v %v\n", bSl, len(bSl), cap(bSl))

	// append - thêm phần tử vào slice
	aSl = append(aSl, 1)
	aSl = append(aSl, 2)
	aSl = append(aSl, 3)
	aSl = append(aSl, 4)
	aSl = append(aSl, 5, 6, 7, 8, 9, 10)
	fmt.Printf("%v %v %v\n", aSl, len(aSl), cap(aSl))

	aSl = append(aSl, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("%v %v %v\n", aSl, len(aSl), cap(aSl))

	// Append giá trị Array vào Slice dùng ...
	aSl = append(aSl, []int{12, 13, 14, 15, 16, 17}...)
	fmt.Printf("%v %v %v\n", aSl, len(aSl), cap(aSl))

	// Stack
	aStack := []int{1, 2, 3, 4, 5, 6}
	// Pop giá trị
	bStack := aStack[:5]

	fmt.Println(bStack)

	// Queue
	aQueue := []int{1, 2, 3, 4, 5, 6}
	bQueue := aQueue[1:]

	fmt.Println(bQueue)

	// Control Slice
	// Nếu dùng append để thay đổi phần tử thì sẽ ảnh hưởng tới Slice ban đầu
	// Nếu không muốn thay đổi giá trị Slice ban đầu thì ta dùng Loop
	aControlSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(aControlSlice)
	bControlSlice := append(aControlSlice[:2], aControlSlice[3:]...)
	fmt.Println(aControlSlice)
	fmt.Println(bControlSlice)
}
