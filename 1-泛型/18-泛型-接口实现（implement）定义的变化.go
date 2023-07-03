package main

func main() {
	// 在GO1.18开始，接口实现（implement）的定义自然也发生了变化
	// 当满足以下条件时，我们可以说 类型 T 实现了接口 I (type T implements interface I)
	//条件：T 不是接口时，类型T是接口 I 代表的类型集合中的一个成员 （T is an element of the type set of I）
	//条件：T 是接口时，T 接口代表的类型集是 I 代表的类型集的子集（Type set of T is a subset of the type set of I）
}
