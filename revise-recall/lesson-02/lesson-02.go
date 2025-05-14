package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Khai báo biến trong Go")
	fmt.Println("Có 3 cách khai báo biến trong Go, bao gồm: ")
	fmt.Println("Cách 1: Khai báo biến đầy đủ với từ khoá(keyword)");
	fmt.Println("var ten_bien kieu_du_lieu = gia_tri")
	fmt.Println("var variable_name data_type = value")
	fmt.Println("Ví dụ 1: ")
	fmt.Println("")
	var myFullName string = "Nguyễn Thế Anh";
	var myAge int = 21;
	var isHandsome bool = true;
	fmt.Println(myFullName,myAge,isHandsome)
	fmt.Println("")

	fmt.Println("Cách 2: Khai báo biến ngắn gọn với :=")
	fmt.Println("Ví dụ 2: ")
	mySchool := "FPT";
	myMajor := "Web Programming"
	myGPA := 9.5
	fmt.Println(mySchool,myMajor,myGPA)
	fmt.Println("")

	fmt.Println("Cách 3: Khai báo biến mà không gán giá trị")
	var myType string
	var myPhoneNumber int
	var myIELTS float64
	var isLazy bool
	fmt.Println(myType,myPhoneNumber,myIELTS, isLazy)
	fmt.Println("")

	fmt.Println("Khai báo hằng số trong Go")
	fmt.Println("Có hai cách khai báo hằng số trong Go: ")
	fmt.Println("Cách 1: Khai báo một hằng số đơn lẻ")
	fmt.Println("Ví dụ 1: ")
	const APP_NAME = "\"Golang Tutorials\"";
	const PI = 3.14159
	const MAX_SIZE = 1024
	const STATUS_ERROR = 500
	const STATUS_OK = 200
	fmt.Println("APP_NAME:", APP_NAME)
	fmt.Println("PI bằng:",PI)
	fmt.Println("MAX_SIZE, STATUS_OK, STATUS_ERROR: ", MAX_SIZE, STATUS_OK, STATUS_ERROR)

	fmt.Println("Cách 2: Khai báo nhiều hằng số cùng lúc (hay còn gọi là nhóm const)")
	fmt.Println("Ví dụ 2: ")
	const (
		APP_NAME_1 = "\"Golang Tutorials\""
		MAX_SIZE_1 = 1024
		STATUS_FORBIDDEN = 404
		PI_1 = 3.14
	)
	fmt.Println("APP_NAME_1, MAX_SIZE_1, STATUS_FORBIDEN, PI_1: ", APP_NAME_1, MAX_SIZE_1, STATUS_FORBIDDEN,PI_1)

	fmt.Println("Chỉ sử dụng được iota trong khai báo nhiều hằng số cùng lúc")
	fmt.Println("Ví dụ 3:")
	const (
		MONDAY = iota + 2 
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY = 1
	)
	fmt.Println("Các thứ trong 1 tuần, 1 tuần có 7 ngày từ thứ 2 đến chủ nhật: ", 
	MONDAY,
	TUESDAY,
	WEDNESDAY,
	THURSDAY,
	FRIDAY,
	SATURDAY,
	SUNDAY)
}