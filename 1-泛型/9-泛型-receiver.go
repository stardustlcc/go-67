package main

import "fmt"

type MySlice[T int | float32] []T

func (my MySlice[T]) sum() T {
	var sum T

	for _, value := range my {
		sum += value
	}
	return sum
}

func main() {

	var my MySlice[int] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(my.sum())

	var my1 MySlice[float32] = MySlice[float32]{0.1, 0.2, 0.3, 0.4}
	fmt.Println(my1.sum())

}
