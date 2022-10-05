package main // in go, this is a reserved name for executable files

import "fmt" // part of the go std library (https://pkg.go.dev/fmt@go1.19.2#pkg-variables)

func main() {
	fmt.Println("Hi there!")
}

// go build is the compile command
// go run builds and executes code
// go fmt formats all the code in each of the file in the pwd
// go install compiles and installs a package
// go test runs any tests associated with the current project
