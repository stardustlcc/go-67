package main

import (
	"fmt"
)

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

func main() {
	var a MyMap[string, float64] = map[string]float64{
		"aaa": 8.1,
		"bbb": 9.3,
	}

	//参考连接：https://mp.weixin.qq.com/s/INj6uAuHHJ5xmSilU6XSUA

	fmt.Println(a)
}
