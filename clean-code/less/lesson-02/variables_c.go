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

	fmt.Println("var name string = \"Golang\" ")
	fmt.Println("var age int = 21")              
	fmt.Println("var isActive bool = true")      

	var name string = "Golang" 	
	var age int = 21          	
	var isActive bool = false

	fmt.Println(name)		
	fmt.Println(age)    	
	fmt.Println(isActive)
	fmt.Println("")


	fmt.Println("2. Khai báo biến ngắn gọn với := (chỉ sử dụng trong hàm) (short declaration):")
	fmt.Println("tên_biến := giá_trị")
	fmt.Println("Sử dụng cách này chương trình sẽ tự động ngầm hiểu và xác định kiểu dữ liệu")
	fmt.Println("")

	fmt.Println("Ví dụ:")

	fmt.Println("name1 := \"Golang\" ") 
	fmt.Println("age1 := 22")           
	fmt.Println("isActive1 := false")

	name1 := "Golang" 	
	age1 := 22        	
	isActive1 := false 	

	fmt.Println(name1)     	
	fmt.Println(age1)      	
	fmt.Println(isActive1) 
	fmt.Println("")

	fmt.Println("3. Khai báo biến với từ khóa var (không gán giá trị):")
	fmt.Println("var tên_biến kiểu_dữ_liệu")
	fmt.Println("")

	fmt.Println("Khi khai báo biến mà không gán giá trị, Go sẽ tự động gán giá trị mặc định cho biến đó.")
	fmt.Println("Giá trị mặc định của các kiểu dữ liệu trong Go:")
	fmt.Println("")
	
	fmt.Println("int: 0")
	fmt.Println("float: 0.0")
	fmt.Println("string: \"\" (chuỗi rỗng)")
	fmt.Println("bool: false")
	fmt.Println("")

	fmt.Println("Ví dụ:")
	fmt.Println("var count int")
	fmt.Println("var price float64")
	fmt.Println("var name2 string")
	fmt.Println("var isActive2 bool")

	var count int 		
	var price float64
	var name2 string
	var isActive2 bool
	fmt.Println("")

	fmt.Println(count)     	
	fmt.Println(price)    	
	fmt.Println(name2)     
	fmt.Println(isActive2)
	fmt.Println("")

	fmt.Println("1.2 Khai báo hằng số")
	fmt.Println("Hằng số là một giá trị không thay đổi trong suốt quá trình thực thi chương trình.")
	fmt.Println("Hằng số được khai báo bằng từ khóa const.")
	fmt.Println("Cú pháp khai báo hằng số:")
	fmt.Println("const tên_hằng số = giá_trị")
	fmt.Println("Ví dụ:")

	fmt.Println("const PI = 3.14...")
	fmt.Println("const APP_NAME = \"Golang Tutorial\"")
	const PI = 3.14159
	const APP_NAME = "Golang Tutorial"
	fmt.Println("PI =", PI)
	fmt.Println("APP_NAME =", APP_NAME)
	fmt.Println("")

	fmt.Println("1.3 Khai báo nhiều hằng số: ")
	fmt.Println("Ta khai báo nhiều hằng số khi ta muốn nhóm chung vào một group.")
	fmt.Println("Cú pháp:")
	fmt.Println(`
	const (
		tên_hằng số1 = giá_trị1
		tên_hằng số2 = giá_trị2
		tên_hằng số3 = giá_trị3
	)`);
	
	const (
		STATUS_OK = 200
		STATUS_ERROR = 500
		STATUS_NOT_FOUND = 404
		STATUS_UNAUTHORIZED = 401
		STATUS_FORBIDDEN = 403
		STATUS_BAD_REQUEST = 400
		STATUS_INTERNAL_SERVER_ERROR = 500
		STATUS_SERVICE_UNAVAILABLE = 503
		STATUS_GATEWAY_TIMEOUT = 504
		STATUS_NOT_IMPLEMENTED = 501
		STATUS_CONFLICT = 409
		MAX_SIZE = 1024
		MIN_SIZE = 1
		DEFAULT_TIMEOUT = 30
	)
	fmt.Println("const STATUS_OK =", STATUS_OK)
	fmt.Println("const STATUS_ERROR =", STATUS_ERROR)
	fmt.Println("const STATUS_NOT_FOUND =", STATUS_NOT_FOUND)
	fmt.Println("const STATUS_UNAUTHORIZED =", STATUS_UNAUTHORIZED)
	fmt.Println("const STATUS_FORBIDDEN =", STATUS_FORBIDDEN)
	fmt.Println("const STATUS_BAD_REQUEST =", STATUS_BAD_REQUEST)
	fmt.Println("const STATUS_INTERNAL_SERVER_ERROR =", STATUS_INTERNAL_SERVER_ERROR)
	fmt.Println("const STATUS_SERVICE_UNAVAILABLE =", STATUS_SERVICE_UNAVAILABLE)
	fmt.Println("const STATUS_GATEWAY_TIMEOUT =", STATUS_GATEWAY_TIMEOUT)
	fmt.Println("const STATUS_NOT_IMPLEMENTED =", STATUS_NOT_IMPLEMENTED)
	fmt.Println("const STATUS_CONFLICT =", STATUS_CONFLICT)
	fmt.Println("const MAX_SIZE =", MAX_SIZE)
	fmt.Println("const MIN_SIZE =", MIN_SIZE)
	fmt.Println("const DEFAULT_TIMEOUT =", DEFAULT_TIMEOUT)
	fmt.Println("")
	fmt.Println("-	Sử dụng iota để khai báo hằng số")
	fmt.Println("iota là một bộ đếm được tự động tăng trong các khối khai báo const.")
	fmt.Println("")
	fmt.Println("Ví dụ:")
	fmt.Println(`const (
		MONDAY = iota
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
		// Các giá trị của MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY lần lượt là 0, 1, 2, 3, 4, 5, 6
		// Nếu ta muốn giá trị của MONDAY là 1, thì ta có thể khai báo như sau:
		MONDAY = iota + 1
		// Các giá trị của MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY lần lượt là 1, 2, 3, 4, 5, 6, 7
	)`)
	const (
		MONDAY = iota
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
		// Các giá trị của MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY lần lượt là 0, 1, 2, 3, 4, 5, 6
		// Nếu ta muốn giá trị của MONDAY là 1, thì ta có thể khai báo như sau:
	)
	const (
		MONDAY1 = iota + 1
		TUESDAY2
		WEDNESDAY3
		THURSDAY4
		FRIDAY5
		SATURDAY6
		SUNDAY7	
	)
	fmt.Println("const MONDAY =", MONDAY)
	fmt.Println("const TUESDAY =", TUESDAY)
	fmt.Println("const WEDNESDAY =", WEDNESDAY)
	fmt.Println("const THURSDAY =", THURSDAY)
	fmt.Println("const FRIDAY =", FRIDAY)
	fmt.Println("const SATURDAY =", SATURDAY)
	fmt.Println("const SUNDAY =", SUNDAY)
	fmt.Println("")
	fmt.Println("const MONDAY1 =", MONDAY1)
	fmt.Println("const TUESDAY2 =", TUESDAY2)
	fmt.Println("const WEDNESDAY3 =", WEDNESDAY3)
	fmt.Println("const THURSDAY4 =", THURSDAY4)
	fmt.Println("const FRIDAY5 =", FRIDAY5)
	fmt.Println("const SATURDAY6 =", SATURDAY6)
	fmt.Println("const SUNDAY7 =", SUNDAY7)
	fmt.Println("")
}
