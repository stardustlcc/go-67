package main

import "fmt"

type WowStruct[T int | float32, S []T] struct {
	Data     S //切片类型
	MaxValue T
	MinValue T
}

func main() {
	var myStruct WowStruct[int, []int] = WowStruct[int, []int]{
		Data:     []int{1, 2, 3},
		MaxValue: 10,
		MinValue: 1,
	}
	fmt.Println(myStruct) //{[1 2 3] 10 1}

	var myStructF WowStruct[float32, []float32] = WowStruct[float32, []float32]{
		Data:     []float32{1.1, 1.2},
		MaxValue: 100.1,
		MinValue: 0.01,
	}
	fmt.Println(myStructF) //{[1.1 1.2] 100.1 0.01}
}
