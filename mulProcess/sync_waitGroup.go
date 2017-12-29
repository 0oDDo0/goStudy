package main

import (
	"sync"
	"fmt"
)

func waitGroup()  {
	var wg sync.WaitGroup
	wg.Add(2) // 两次Done, 设置等待的数量, num = 2

	wg.Done() // 每调用一次, num--
	wg.Done()

	wg.Wait() // 等待, 直至num = 0
	fmt.Println("finish")
}

func main() {
	waitGroup()
}