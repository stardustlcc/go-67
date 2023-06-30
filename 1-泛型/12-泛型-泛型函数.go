package main

import "fmt"

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

func main() {
	num := Add[int](100, 200)
	fmt.Println(num)

	//自动推导出类型实参，实际上传入实参步骤还是发生了
	num1 := Add(1, 2)
	fmt.Println(num1)

	fnum := Add[float32](1.1, 1.8)
	fmt.Println(fnum)

	fvNum := Add[float64](100.1, 200.2)
	fmt.Println(fvNum)
}
