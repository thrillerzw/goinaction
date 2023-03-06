package main

import (
	"fmt"
	"sync"
)

// 类似java CountDownLatch
var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		//按照格式化文本输出内容到os.stdout。
		fmt.Printf("Received %d ", i)
	}
	//WaitGroup减1
	wg.Done()
}

// 输出结果：Received 1 Received 2 Received 3 Received 4 Received 5 Received 6 Received 7 Received 8 Received 9 Received 10
// main is the entry point for the program.
func main() {
	//创建无缓冲通道
	c := make(chan int)
	//多线程，先加入调度器全局队列，之后调度器会分配给一个逻辑处理器，并放到逻辑处理器本地运行队列。
	go printer(c)
	//计数加1，表示要等等1个goroutine
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	// 等待goroutine结束
	wg.Wait()
}
