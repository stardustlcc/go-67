package main

import "fmt"

// 目前，方法不支持泛型
type A struct {
}

//错误，不支持泛型方法
// func (receiver A) Add[T int|float32](a T, b T) T {
// 	return a + b
// }

type B[T int | float32] struct{}

// 但是方法可以使用类型定义的形参T
func (receiver B[T]) Add(a T, b T) T {
	return a + b
}

func main() {
	var b B[int]
	sum := b.Add(100, 200)
	fmt.Println(sum)

	var c B[float32]
	sum1 := c.Add(1.1, 2.6)
	fmt.Println(sum1)
}
