package main

import (
	"fmt"
	"sync"
)

/*
✅Channel
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
var wg sync.WaitGroup

func test12() {
	wg.Add(2)

	ch := make(chan int, 100)
	go producer(ch)
	go consumer(ch)

	wg.Wait()
}

func producer(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch chan int) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}
