package main

import "fmt"

func main() {
	testCase := struct {
		Name string
		Age  int
	}{
		Name: "cici",
		Age:  18,
	}
	fmt.Println(testCase)

	//下面的案例是错误的，会报错
	// testCase1 := struct[T int| string]{
	// 	Name T
	// 	Age int
	// }
}
