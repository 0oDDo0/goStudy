package main

import (
	"sync"
	"time"
)

// 读写锁
var m = new(sync.RWMutex)
func read(i int) {
	println(i, "read start")

	// 当写锁没锁时, 可多个协程读
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read end")
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()

	println(i, "write end")
}
func main()  {
	m = new(sync.RWMutex)
	// 写的时候, 读锁等着
	go write(1)
	go read(2)
	go read(2)
	go read(2)
	go write(3)
	time.Sleep(4 * time.Second)
}