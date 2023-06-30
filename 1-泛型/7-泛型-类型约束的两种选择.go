package main

// 结构体 WowStructT 和 结构体 WowStructF 的功能实现差不多
// 第一种写法
type WowStructT[T int | string] struct {
	Name string
	Data []T
}

// 第二种写法
type WowStructF[T []int | []string] struct {
	Name string
	Data T
}

// 但是像这种定义，需要复用同一个类型时，建议使用第一种写法
type WowStructS[T int | string] struct {
	Name    string
	Data    []T
	MinData T
}

func main() {

}
