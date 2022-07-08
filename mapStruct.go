package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Map - tập hợp các ánh xạ giữa các khóa gọi là Key tham chiếu đến giá trị của nó là Value
	// tất cả kiểu dữ liệu primitive đều có thể làm Key trừ kiểu Slice
	studentNameAgeMap := map[string]int32{
		"Hieu": 22,
		"Tom":  35,
		"Mike": 44,
	}

	fmt.Printf("%v %T\n", studentNameAgeMap, studentNameAgeMap)

	// make - khai báo map khi chưa biết key và value
	valueMap := make(map[string]int)

	valueMap = map[string]int{
		"HAHA": 20,
		"HEHE": 30,
		"HIHI": 40,
	}
	fmt.Printf("%v %T\n", valueMap, valueMap)

	// Truy xuất map
	fmt.Printf("%v %T\n", valueMap["HAHA"], valueMap["HAHA"])

	valueMap["HOHO"] = 50
	fmt.Printf("%v %T\n", valueMap, valueMap)

	// Edit
	valueMap["HEHE"] = 100
	fmt.Printf("%v %T\n", valueMap, valueMap)

	// Delete
	delete(valueMap, "HIHI")
	fmt.Printf("%v %T\n", valueMap, valueMap)

	// Check xem Key đã bị xóa hay chưa
	valueHIHI, isExist := valueMap["HIHI"]
	fmt.Printf("%v %T\n", valueHIHI, valueHIHI)
	fmt.Printf("%v %T\n", isExist, isExist)

	_, isExistHIHI := valueMap["HIHI"]
	fmt.Printf("%v %T\n", isExistHIHI, isExistHIHI)

	// Len - độ dài Map
	fmt.Println(len(valueMap))

	// Copy Map - tương tự như Slice đều là tham chiếu đến địa chỉ con trỏ.
	copyMap := valueMap
	copyMap["HOHO"] = 99

	fmt.Printf("%v %T\n", valueMap, valueMap)
	fmt.Printf("%v %T\n", copyMap, copyMap)

	// Struct - tập hợp các kiểu dữ liệu khác mà ta tự định nghĩa

	type Student struct {
		number   int
		name     string
		isMale   bool
		subjects []string
	}

	// Declare Struct

	studentOne := Student{
		number: 3,
		name:   "David",
		isMale: true,
		subjects: []string{
			"Math",
			"English",
			"Computer",
		},
	}

	fmt.Println(studentOne)

	var studentTwo Student = Student{
		4,
		"Anna",
		false,
		[]string{
			"Draw",
			"Photoshop",
			"Literature",
		},
	}

	// Truy xuất Struct
	fmt.Println(studentTwo)

	// Truy xuất 1 thuộc tính
	fmt.Println(studentTwo.name)

	// Khai báo đối tượng rỗng
	studentEmpty := Student{}
	studentEmpty.name = "Alo"
	studentEmpty.isMale = true

	fmt.Println(studentEmpty)

	// Khai báo inline struct - vừa khai báo struct vừa tạo đối tượng

	studentInline := struct {
		name string
	}{
		name: "Hieu",
	}

	fmt.Println(studentInline)

	// Copy value - không thay đổi thuộc tính ban đầu
	copyStudent := studentInline

	copyStudent.name = "Vahalla"

	fmt.Println(studentInline)
	fmt.Println(copyStudent)

	// Copy value trỏ đến vùng nhớ
	// Xài con trỏ để trỏ đến vùng nhớ mà không cần tạo mới để tiết kiệm bộ nhớ.
	copyStudentPointer := &studentInline

	copyStudentPointer.name = "Vahalla"

	fmt.Println(studentInline)
	fmt.Println(*copyStudentPointer)

	// tích hợp struct trong struct

	// Struct nền tảng
	type People struct {
		name string
		age  int
	}

	// sử dụng struct People - tương tự tính thừa kế
	type StudentP struct {
		People
		number   int
		subjects []string
	}

	studentEmbbeded := StudentP{}
	studentEmbbeded.name = "Hieu"
	studentEmbbeded.age = 22
	studentEmbbeded.number = 9
	studentEmbbeded.subjects = []string{"Math", "English", "Computer"}

	fmt.Printf("%v %T \n", studentEmbbeded, studentEmbbeded)

	// Tag - `` thuộc tính trong field number

	type StudentP2 struct {
		People
		number   int `Maximum can't over 100`
		subjects []string
	}

	studentTag := StudentP2{}
	t := reflect.TypeOf(studentTag)
	field, _ := t.FieldByName("number")
	fmt.Println(field)
}
