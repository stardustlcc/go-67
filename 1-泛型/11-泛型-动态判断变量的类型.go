package main

import (
	"fmt"
	"reflect"
)

type Queue[T interface{}] struct {
	Elem []T
}

func (q Queue[T]) Put(value T) {
	// printf() 输出变量的类型，（底层通过反射实现的）
	fmt.Printf("%T--\n", value)

	//泛型类型的形参是不可以使用类型断言的
	//value.(type)

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	}
}

func main() {

	// 原始的类型断言案例
	var i interface{} = 100
	var s interface{} = "name"
	var b interface{} = true

	check(b)
	check(s)
	check(i)

	// 泛型类型不支持类型断言
	// 但是可以通过反射实现断言。但是就违背了最初的原则。
	// 为了避免反射而选择而选择了泛型，结果现在为了一些功能在泛型中使用反射，此时应该思考是否需要使用泛型
	var q Queue[int]
	q.Put(1)
	var str Queue[string]
	str.Put("aaa")
}

func check(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	}
}
