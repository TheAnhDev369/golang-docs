package main

import "fmt"

func main() {
	//	Hello, World! from Go
	fmt.Println("Hello, World!")

	//	Chạy các lệnh Go cơ bản
	//	Những lệnh này chạy độc lập, không cần go mod init, không phụ thuộc vào thư viện ngoài
	fmt.Println("Check version:  		go version")							//	Hiển thị phiên bản Go hiện tại
	fmt.Println("Run Go Program: 		go run filename.go");					//	Chạy file Go trực tiếp
	fmt.Println("Check Go Environment: 	go env")								//	Hiển thị các biến, thông tin môi trường Go hiện tại
	fmt.Println("Check Go Build: 		go build filename.go")					//	Xây dựng file Go thành file thực thi
	fmt.Println("Check Go Clean: 		go clean -cache -testcache -modcache")	//	Xóa cache, bộ nhớ tạm thời của Go
	fmt.Println("Check Go Doc: 			go doc fmt")							//	Xem tài liệu của package fmt	
	fmt.Println("Check Go Fmt: 			go fmt")								//	Định dạng lại code Go


	//	Chạy các lệnh Go liên quan đến package
	//	Các lệnh này cần phải có go mod init, go mod tidy trước đó
	//	Để tải về các package cần thiết cho project Go
	// 	Những lệnh cần CÓ module (go.mod) mới chạy được ổn định:
	fmt.Println("Check Go Install: 		go install")							//	Cài đặt package Go vào thư mục $GOPATH/bin
	fmt.Println("Check Go Modules: 		go mod tidy")							//	Tải về các package cần thiết cho project Go
	fmt.Println("Check Go Get: 			go get github.com/user/repo")			//	Tải về package Go từ github
	fmt.Println("Check Go Test: 		go test -v")							//	Chạy các test case trong file test.go
	fmt.Println("Check Go Vet: 			go vet")								//	Kiểm tra lỗi trong code Go						
}


/** package main
 * Giải nghĩa: 
 * Đây là cách Go khai báo một package (gói) – giống như “thư mục logic” chứa các file code.
 * main là package đặc biệt vì nó đánh dấu: đây là chương trình chính có thể chạy được (chứ không phải thư viện).
 * Nếu bạn đặt là package utils thì Go sẽ hiểu đây là 1 thư viện phụ chứ không chạy được trực tiếp
 
 * => Nếu bạn không có package main thì Go sẽ báo lỗi: no main function
 * => Muốn chạy được chương trình Go ⇒ phải có package main.
 */

/** import "fmt"
 * Giải nghĩa:
 * import dùng để nhúng (import) một package có sẵn trong thư viện chuẩn của Go.
 * "fmt" là viết tắt của format, chứa các hàm để xuất thông tin ra màn hình, ví dụ: Print, Println, Printf,...
 * 
 * => Nhờ import "fmt" bạn mới dùng được fmt.Println(...) bên dưới.
 * => Nếu bạn không import "fmt" thì Go sẽ báo lỗi: undefined: fmt
 */

/** func main() {}
 * Giải nghĩa:
 * func là từ khóa dùng để khai báo một hàm (function) hay từ khóa định nghĩa hàm.
 * main là tên hàm, là hàm chính của chương trình.
 * Hàm main() là hàm đầu tiên được chạy khi bạn chạy chương trình Go.
 * main() là hàm đặc biệt mà Go tự động gọi đầu tiên khi bạn chạy chương trình
 * () là dấu ngoặc đơn, dùng để chứa các tham số truyền vào hàm (nếu có).
 * {} là dấu ngoặc nhọn, dùng để chứa nội dung bên trong hàm.
 *
 * => Không có func main(), chương trình Go sẽ không biết chạy từ đâu.
 */

 /** fmt.Println("Hello, World!");
 * Giải nghĩa:
 * fmt.Println() là hàm in ra dòng chữ kèm dấu xuống dòng (\n)
 * "Hello, World!" là chuỗi (string) bạn muốn in.
 * Println viết tắt của Print Line, có nghĩa là in ra dòng chữ và xuống dòng.
 * Nếu bạn không có Println thì nó sẽ không xuống dòng.
 * Nếu bạn muốn in mà không xuống dòng thì dùng Print.
 * => fmt.Println() dùng để xuất nội dung ra màn hình một cách rõ ràng và có dấu xuống dòng.
*/
