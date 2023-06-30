package main

import "fmt"

// 先定义一个泛型类型 Slice[]
type Slice[T int | string | bool | float32 | float64] []T

// 错误，因为 Slice中不包含 uint 类型
//type UintSlice[T uint | uint8] Slice[T]

// 正确，根据泛型类型的 Slice 定义新的泛型
type MergeSlice[T int | string] Slice[T]

// 正确，基于 MergeSlice 类型定义新的泛型类型
type IntSlice[T int] MergeSlice[T]

// 在 map 中套一个泛型类型 Slice[T]
type WowMap[T int | string] map[string]Slice[T]

// 在 map 中套 Slice 的另一种写法
type WowMap2[T Slice[int] | Slice[string]] map[string]T

func main() {
	//套娃后的 map 使用
	var myMap WowMap2[Slice[int]] = WowMap2[Slice[int]]{
		"myslice": Slice[int]{1, 2, 3, 4},
	}

	fmt.Println(myMap)
}
