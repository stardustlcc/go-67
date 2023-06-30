package main

import "fmt"

// 泛型约束列表， 因为太长了不好管理
type Slice[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64] []T

// 接口，通过接口来管理泛型约束列表
type IntUintFloat interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// 使用接口定义 泛型类型
type SliceInter[T IntUintFloat] []T

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

// 貌似这样定义，参数还是很长
type SliceT[T Int | Uint | Float] []T

// 继续优化
type AFace interface {
	Int | Uint | Float
}

type SliceF[T AFace] []T

func main() {

	var a SliceF[int] = SliceF[int]{
		1, 2, 3,
	}
	fmt.Println(a)
}
