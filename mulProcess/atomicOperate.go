package main

import (
	"sync/atomic" // 提供原子操作的库函数
	"fmt"
	"os/exec"
)
/*
atomic operation 原子操作, 执行过程中不能被中断的操作, 原子操作必须由一个单一的汇编指令表示
<所有的系统调用都是原子操作>
*/
func testAtomic(){
	var a int32 = 0
	atomic.AddInt32(&a, 1) // 原子操作的加
	atomic.
	fmt.Println(a)
}

func pipe() string {
	cmd := exec.Command("python", "--version")
	outPipe, err := cmd.StdoutPipe()
	if err != nil { fmt.Println(err); return "output error" }
	if err := cmd.Start(); err != nil { fmt.Println(err); return "start error " }
	out := make([]byte, 30)
	if _, err := outPipe.Read(out); err != nil { fmt.Println(err); return "read error" }
	return string(out)
}

func main()  {
	testAtomic()
	fmt.Println(pipe())
}