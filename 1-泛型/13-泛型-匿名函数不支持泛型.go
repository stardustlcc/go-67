package main

import "fmt"

func main() {
	fn := func(a, b int) int {
		return a + b
	}

	num := fn(1, 2)
	fmt.Println(num)

	//匿名函数不能自定义类型实参
	// fn1 := func[T int| float32](a T, b T) T {
	// 	return a + b
	// }

	//但是匿名函数可以使用别处定义好的类型实参
	a := MyFunc[int](100, 10)
	fmt.Println(a)
}

func MyFunc[T int | float32 | float64](a, b T) T {
	fu1 := func(i T, j T) T {
		return i + j
	}
	return fu1(a, b)
}
