package main

import (
	"fmt"
)

type Wow[T int | string] int

func main() {
	var age Wow[int] = 100
	fmt.Println(age)

	var name Wow[string] = 111
	fmt.Println(name)

	//虽然使用的形参定义了类型，但是因为底层是 int ，则始终不能将字符串传递
	//var address Wow[string] = "CIC"
}
