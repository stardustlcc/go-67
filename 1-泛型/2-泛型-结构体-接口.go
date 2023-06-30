package main

import (
	"fmt"
)

// 结构体
type MyStruct[T int | string] struct {
	Name string
	Data T
}

// 结构
type PrintData[T int | float32 | float64] interface {
	Print(data T)
}

// 通道
type MyChan[T int | string] chan T

// 结构体实现接口
func (my *MyStruct[T]) Print(data T) {
	fmt.Println(data)
}

func main() {

	//结构体使用
	var myStruct MyStruct[int] = MyStruct[int]{
		Name: "cici",
		Data: 100,
	}

	myStruct.Print(myStruct.Data) //100

	var myStructStr MyStruct[string] = MyStruct[string]{
		Name: "cici",
		Data: "上海",
	}
	myStructStr.Print(myStructStr.Data) //上海

	//chan 的使用
	var myChan MyChan[int] = make(MyChan[int], 1)
	myChan <- 100
	fmt.Println(<-myChan) //100

	var myChanStr MyChan[string] = make(MyChan[string], 1)
	myChanStr <- "你好啊"
	fmt.Println(<-myChanStr) //你好啊

}
