package main

import "fmt"

//队列：先入先出的数据结构，数据只能从队尾放入，从队首取出。先放入的数据优先被取出

// 定义一个空接口结构体，可以使用任何类型来实例化此泛型 Queue[T]
type Queue[T interface{}] struct {
	elem []T
}

// 将数据放入队列尾部
func (q *Queue[T]) Put(value T) {
	q.elem = append(q.elem, value)
}

// 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elem) == 0 {
		return value, true
	}
	value = q.elem[0]
	q.elem = q.elem[1:]
	return value, len(q.elem) == 0
}

// 队列大小
func (q Queue[T]) Size() int {
	return len(q.elem)
}

func main() {
	var q Queue[int]
	q.Put(1)
	q.Put(2)
	q.Put(3)
	q.Put(4)
	q.Put(5)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	var qStr Queue[string]
	qStr.Put("cici")
	qStr.Put("age")
	qStr.Put("addrss")

	fmt.Println(qStr.Size())
	fmt.Println(qStr.Pop())
	fmt.Println(qStr.Pop())
	fmt.Println(qStr.Pop())
}
