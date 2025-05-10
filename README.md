# golang-docs
# Giới thiệu về ngôn ngữ lập trình Go (Golang)

##  1. Tổng quan về Golang

**Go**, còn gọi là **Golang**, là một ngôn ngữ lập trình mã nguồn mở được phát triển bởi Google vào năm 2007 và công bố ra công chúng vào năm 2009. Go được thiết kế để đơn giản, hiệu quả, và dễ sử dụng trong các hệ thống lớn, đặc biệt là phát triển phần mềm backend, công cụ mạng, và hệ thống phân tán.

>  Tác giả chính: Robert Griesemer, Rob Pike, và Ken Thompson.

##  2. Tính năng nổi bật

- **Hiệu suất cao**: Go được biên dịch thành mã máy, gần với hiệu năng của C/C++.
- **Quản lý bộ nhớ tốt**: Hệ thống quản lý bộ nhớ tự động (garbage collection) mạnh mẽ.
- **Đơn giản và dễ học**: Cú pháp rõ ràng, dễ đọc, dễ duy trì.
- **Concurrency mạnh mẽ**: Thư viện `goroutines` và `channels` giúp xử lý song song dễ dàng.
- **Hệ thống công cụ tích hợp**: Bao gồm `go build`, `go test`, `go fmt`, v.v.
- **Biên dịch nhanh**: Thời gian biên dịch cực kỳ nhanh, rất thích hợp cho phát triển nhanh và CI/CD.

##  3. Cú pháp cơ bản

```go
package main

import "fmt"

func main() {
    fmt.Println("Xin chào, Golang!")
}
```

##  4. Ứng dụng phổ biến
- Viết **microservices** và API backend (gRPC, REST).
- Phát triển **hệ thống phân tán** và cloud-native (Kubernetes viết bằng Go).
- Viết **CLI tools** và ứng dụng DevOps.
- **Blockchain**, ví dụ: Ethereum client (Geth) dùng Go.
- Xây dựng **web frameworks** (Gin, Echo, Fiber...).

##  5. Cài đặt Go trên Windows

### Bước 1: Tải về
Tải Golang từ trang chính thức:  
🔗 https://go.dev/dl/

### Bước 2: Cài đặt
- Mở file `.msi` và thực hiện theo hướng dẫn.
- Mặc định sẽ cài vào: `C:\\Program Files\\Go`.

### Bước 3: Cấu hình môi trường
- Thêm đường dẫn `C:\\Program Files\\Go\\bin` vào biến môi trường `PATH`.

Kiểm tra bằng lệnh:

```bash
go version
```

##  6. Cấu trúc một project Go
```bash
my-project/
│
├── go.mod           # Tệp quản lý module
├── main.go          # Tập tin chính
├── utils/
│   └── helper.go    # Các tiện ích
```
### Tạo module:

```bash
go mod init my-project
```

### Chạy chương trình:

```bash
go run main.go
```

##  7.  Các framework và thư viện phổ biến
| Tên    | Mô tả                           |
|--------|----------------------------------|
| Gin    | Web framework nhanh, nhẹ         |
| Echo   | Web framework dễ mở rộng         |
| Cobra  | Tạo CLI tool                     |
| GORM   | ORM cho kết nối cơ sở dữ liệu    |
| Viper  | Quản lý config mạnh mẽ           |

##  8. Cộng đồng và học tập
-   Trang chủ: https://go.dev/

-   Tài liệu chính thức: https://pkg.go.dev/

-   Khóa học Go miễn phí: https://tour.golang.org/

-   Cộng đồng VN: Facebook group "Golang Việt Nam", nhóm Telegram...