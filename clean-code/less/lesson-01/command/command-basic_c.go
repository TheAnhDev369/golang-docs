package main

import "fmt"

func main() {
	fmt.Println("Check version:  		go version")							
	fmt.Println("Run Go Program: 		go run filename.go");					
	fmt.Println("Check Go Environment: 	go env")								
	fmt.Println("Check Go Build: 		go build filename.go")					
	fmt.Println("Check Go Clean: 		go clean -cache -testcache -modcache")	
	fmt.Println("Check Go Doc: 			go doc fmt")								
	fmt.Println("Check Go Fmt: 			go fmt")								


	fmt.Println("Check Go Install: 		go install")							
	fmt.Println("Check Go Modules: 		go mod tidy")							
	fmt.Println("Check Go Get: 			go get github.com/user/repo")			
	fmt.Println("Check Go Test: 		go test -v")							
	fmt.Println("Check Go Vet: 			go vet")														
}
