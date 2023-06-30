package main

import (
	"fmt"
	"time"
)

func main() {

	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int

	fmt.Println(ch1, ch2, ch3) //nil
	fmt.Println("----------------------")

	go func() {
		fmt.Println(<-ch1)
	}()
	//无缓存通通道，得有人接收，不然会阻塞
	ch1 = make(chan int)
	ch2 = make(chan bool)
	ch3 = make(chan []int)
	fmt.Println(ch1, ch2, ch3)
	ch1 <- 1
	fmt.Println("----------------------")

	//有缓存通道，可不用前置有人接收
	ch4 := make(chan int, 10)
	ch4 <- 10

	fmt.Println(<-ch4)
	fmt.Println("----------------------")

	//多返回值模式
	ch5 := make(chan string, 1)
	ch5 <- "hello world"
	close(ch5)
	value, ok := <-ch5

	fmt.Println(value, ok)
	if !ok {
		fmt.Println("通道已关闭")
	}
	fmt.Println(value)
	fmt.Println("----------------------")

	//通过 for range 接收值
	ch1 = make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	go chanPrint(ch1)
	time.Sleep(time.Second * 3)
	fmt.Println("----------------------")

	//单向通道
	ch := product()
	res := consumer(ch)
	fmt.Println(res)
	fmt.Println("----------------------")
}

func product() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func chanPrint(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
