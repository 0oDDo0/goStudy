package main

import (
	"sync"
	"time"
	"fmt"
)

func thread()  {
	type intChan chan int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(intChan, 1)

	thread1 := func(waitsecond time.Duration) {
		time.Sleep(waitsecond * time.Second)
		ch <- 0
		wg.Done()
	}
	thread2 := func() {
		start := time.Now()
		fmt.Println("waiting channel")
		a := <- ch
		fmt.Println("wait time =", time.Now().Sub(start), "a = ", a)
		wg.Done()
	}
	go thread1(3)
	go thread2()
	wg.Wait()
	fmt.Println("func finish")
}

func main()  {
	
}