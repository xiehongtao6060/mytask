package main

import (
	"fmt"
	"sync"
)

/*
✅Goroutine
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func test10() {
	// 创建一个WaitGroup来同步协程
	var wg sync.WaitGroup
	wg.Add(2) // 设置计数器为2，因为有两个协程

	// 第一个协程：打印奇数
	go func() {
		defer wg.Done() // 协程完成时减少计数器
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数:", i)
		}
	}()

	// 第二个协程：打印偶数
	go func() {
		defer wg.Done() // 协程完成时减少计数器
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数:", i)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}
