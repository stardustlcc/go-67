package main

import "fmt"

type AFace interface {
	int | float32
}

type A[T AFace] struct {
	data T
}

type AFace1 interface {
	~int | ~float32
}

type A1[T AFace1] struct {
	data T
}

type myInt int

//限制，使用 ~ 时有一定的限制
//1) ~ 后面的类型不能是接口
//2) ~ 后面的类型必须是基本类型

type _ interface {
	~[]byte
	~int
	//~ error //错误，
}

func main() {

	var a A[int] = A[int]{1}
	fmt.Println(a)

	//错误 除非是加上 ~  ，例如：type A[T ~int | float32] struct
	//var b A[myInt]
	//fmt.Println(b)

	//正确
	var b A1[myInt]
	fmt.Println(b)

}
