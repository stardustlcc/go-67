package main

import "fmt"

// 正确是声明方式
type myType[T int | string] []T

// 错误的声明方式 不能只有类型形参
// type myErrorType[T float32 | int | string] T

// 这种错误声明方式 因为 编译器认为它是一个表达式
// type myErrorType[T *int | *string] []T
// type myErrorType[T(int) | (string)][]T

// 如果想解决上面问题，可以给类型约束包上 interface{} 或加逗号消除歧义

type myStruct[T *int,] []T
type myStruct1[T *float32 | *string,] []T
type myStruct2[T interface{ *int }] []T

func main() {
	var num = 100

	var a myStruct2[*int] = myStruct2[*int]{
		&num,
	}
	fmt.Println(*a[0]) //100
}
