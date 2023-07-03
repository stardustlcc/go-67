package main

import "fmt"

// 在 GO1.18之前，接口（interface）的定义是 一个方法集
// 方法的集合都汇集在了接口中，换一个概念方法集合也可以理解成一个类型集
// 在go1.18开始将接口的定义正式更改为 类型集（Type set）
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 例如
type FloatF interface {
	~float32 | ~float64
}

type SliceT[T FloatF] []T

//如上面的案例，代表了一个类型集合，接口类型的 FloatF 代表了一个类型集合
//所有以float32 或 float64 为底层类型的类型 都在这一类型集之中

func main() {
	var s SliceT[float32] = SliceT[float32]{
		1.1,
		1.2,
	}
	fmt.Println(s)
}
