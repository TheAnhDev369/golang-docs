package main

import (
	"fmt"
)

func main() {
	fmt.Println("1.  Khai báo biến và hằng số trong Go")
	fmt.Println("1.1 Khai báo biến")
	fmt.Println("")

	fmt.Println("Biến là một vùng nhớ có tên, được sử dụng để lưu trữ dữ liệu trong chương trình. \nBiến có thể thay đổi giá trị trong quá trình thực thi chương trình.")
	fmt.Println("")

	fmt.Println("Cú pháp khai báo biến:")
	fmt.Println("Golang cung cấp nhiều cách để khai báo biến, bao gồm:")
	fmt.Println("")

	fmt.Println("1. Khai báo biến đầy đủ với từ khóa var:")
	fmt.Println("var tên_biến kiểu_dữ_liệu = giá_trị")
	fmt.Println("")
	fmt.Println("Ví dụ:")
	//	Ví dụ 
	fmt.Println("var name string = \"Golang\" ")
	fmt.Println("var age int = 21")              
	fmt.Println("var isActive bool = true")      
	
	//	Thực hành
	var name string = "Golang" 	// Khai báo biến name kiểu string và gán giá trị là "Golang"
	var age int = 21          	// Khai báo biến age kiểu int và gán giá trị là 21
	var isActive bool = false	// Khai báo biến isActive kiểu bool và gán giá trị là true
	//	In ra
	fmt.Println(name)		// In ra giá trị của biến name
	fmt.Println(age)    	// In ra giá trị của biến age
	fmt.Println(isActive)	// In ra giá trị của biến isActive
	fmt.Println("")


	fmt.Println("2. Khai báo biến ngắn gọn với := (chỉ sử dụng trong hàm) (short declaration):")
	fmt.Println("tên_biến := giá_trị")
	fmt.Println("Sử dụng cách này chương trình sẽ tự động ngầm hiểu và xác định kiểu dữ liệu")
	fmt.Println("")
	fmt.Println("Ví dụ:")
	//	Ví dụ
	fmt.Println("name1 := \"Golang\" ") 
	fmt.Println("age1 := 22")           
	fmt.Println("isActive1 := false")
	//	Thực hành
	name1 := "Golang" 	// Khai báo biến name kiểu string và gán giá trị là "Golang"
	age1 := 22        	// Khai báo biến age kiểu int và gán giá trị là 22
	isActive1 := false 	// Khai báo biến isActive kiểu bool và gán giá trị là false
	//	In ra
	fmt.Println(name1)     	// In ra giá trị của biến name1
	fmt.Println(age1)      	// In ra giá trị của biến age1
	fmt.Println(isActive1) 	// In ra giá trị của biến isActive1
	fmt.Println("")

	fmt.Println("3. Khai báo biến với từ khóa var (không gán giá trị):")
	fmt.Println("var tên_biến kiểu_dữ_liệu")
	fmt.Println("Khi khai báo biến mà không gán giá trị, Go sẽ tự động gán giá trị mặc định cho biến đó.")
	fmt.Println("Giá trị mặc định của các kiểu dữ liệu trong Go:")
	fmt.Println("int: 0")
	fmt.Println("float: 0.0")
	fmt.Println("string: \"\" (chuỗi rỗng)")
	fmt.Println("bool: false")
	fmt.Println("")
	fmt.Println("Ví dụ:")
	//	Ví dụ
	fmt.Println("var count int")
	fmt.Println("var price float64")
	fmt.Println("var name2 string")
	fmt.Println("var isActive2 bool")
	//	Thực hành
	// 	Khai báo biến count kiểu int và không gán giá trị
	//	Biến count sẽ được gán giá trị mặc định là 0
	var count int 		
	//	Biến price sẽ được gán giá trị mặc định là 0.0
	var price float64
	//	Biến name2 sẽ được gán giá trị mặc định là "" (chuỗi rỗng)
	var name2 string
	//	Biến isActive2 sẽ được gán giá trị mặc định là false
	var isActive2 bool
	//	In ra
	fmt.Println(count)     	// In ra giá trị của biến count
	fmt.Println(price)    	// In ra giá trị của biến price
	fmt.Println(name2)     	// In ra giá trị của biến name2
	fmt.Println(isActive2) 	// In ra giá trị của biến isActive2
	fmt.Println("")
}
