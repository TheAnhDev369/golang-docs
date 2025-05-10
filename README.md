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
