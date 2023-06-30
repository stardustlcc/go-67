package main

import "fmt"

// 原始的切片类型
type IntSlice []int

// 泛型切片
type Slice[T int | float32 | float64 | string] []T

// 泛型数组
type Arr[T string | int] [3]T

// 泛型MAP
type MyMap[KEY int | string, VALUE int | string] map[KEY]VALUE

func main() {

	//原始切片int
	var a IntSlice = []int{1, 2, 3}
	fmt.Println(a)

	//泛型的使用
	var a1 Slice[int] = []int{1, 2, 3}
	fmt.Println(a1)

	var a2 Slice[float32] = []float32{111.33}
	fmt.Println(a2)

	var a3 Slice[float64] = []float64{111212.11121}
	fmt.Println(a3)

	var a4 Slice[string] = []string{"aaa", "bbb", "ccc"}
	fmt.Println(a4)

	//泛型数组的使用
	var arr1 Arr[string] = [3]string{"AAA", "BBB", "CCC"}
	fmt.Println(arr1)

	var arr2 Arr[int] = [3]int{111, 222, 333}
	fmt.Println(arr2)

	//泛型map的使用
	var myMap MyMap[int, string] = make(MyMap[int, string])
	myMap[1] = "aaa"
	myMap[2] = "bbb"
	myMap[3] = "ccc"
	fmt.Println(myMap)

	var myMap1 MyMap[string, string] = MyMap[string, string]{
		"name": "cici",
		"age":  "18",
		"city": "shanghai",
	}
	myMap1["address"] = "上海松江"
	fmt.Println(myMap1)

}

func test(a int, b int) int {
	return a + b
}
