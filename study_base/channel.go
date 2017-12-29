package main

import (
)
import (
	"fmt"
	"time"
	"sync"
)

/*
goroutine(协程), https://studygolang.com/articles/2027
go提倡以通信的手段来共享内存, 即channel(信道, 通道), 它为 "引用" 类型
chan初始化之前值为nil, 用来传递值, 而不是存储值
chan分为以下几种
	chan   read-write
	<-chan only read
	chan<- only write
init: make(chan T, size=0), size相当于channel的容量, 同一时刻通道可以缓冲size个元素,
                            当size等于0时,说明没有缓冲队列, 必须被取出, 不然会造成死锁

var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0)
var ch = make(chan int,10) //有缓冲channel,缓冲大小是10
无缓冲channel在读和写是都会阻塞, 而有缓冲channel在向channel中存入数据没有达到channel缓存总数时, 可以一直向里面存,
直到缓存已满才阻塞. 由于阻塞的存在, 所以使用channel时特别注意使用方法, 防止死锁的产生
func Afuntion(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	ch := make(chan int) //无缓冲的channel
	go Afuntion(ch)
	ch <- 1
	// 输出结果： finish
}
首先创建一个无缓冲ch,然后执行go Afuntion(ch),此时执行＜-ch，则Afuntion这个函数便会阻塞, 不再继续往下执行,
直到主进程中ch<-1向ch中注入数据才解除Afuntion该协程的阻塞．

func main() {
	ch := make(chan int) //无缓冲的channel
	//只是把这两行的代码顺序对调一下
	ch <- 1
	go Afuntion(ch)
	// 输出结果：
	// 死锁，无结果
}
在主协程向ch 中通过ch<-1命令写入数据, 此时主协程阻塞, 就无法执行下面的go Afuntions(ch),无法解除主协程的阻塞状态,则系统死锁

总结:
无缓存的channel:
	放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中, 防止死锁的发生;
	先执行接收数据的协程, 然后在执行放数据的协程, 不然会造成死锁
带缓存channel:
	对于带缓存channel, 只要channel中缓存不满, 则可以一直向 channel中存入数据, 直到缓存已满;
	同理只要channel中缓存不为０, 便可以一直从channel中向外取数据, 直到channel缓存变为０才会阻塞．

相对于不带缓存channel, 带缓存channel不易造成死锁, 可以同时在一个goroutine中放心使用.

close()
close主要用来关闭channel通道, 用法为close(channel), 并且要在生产者的地方关闭channel, 而不是在消费者的地方关闭．
并且关闭channel后, 便不可再向channel中继续存入数据, 但是可以继续从channel中读取数据


特性:
1) goroutine之间传递数据和同步的主要方法
2) 对通道的操作本身也是同步的,即同一时刻, 仅有一个goroutine能向同一个通道发送值, 同时也仅有一个goroutine能从通道接收值
3) channel是一个先进先出的消息队列
4) 通道中的元素具有原子性, 通道中的元素金能被一个goroutine接收, 已被接收的元素值会立刻从通道中删除
5) 当channel满的时候也可以向里面发送数据(当有协程从ch里面取, 不会死锁), 只不过会被阻塞住.


*/
type intChan chan int

func size_0()  {
	// ch0 := make(chan int, 0)  // 编译正确, 但是执行error
	ch0 := make(chan int, 1)  // right
	ch0 <- 1 // size=0时, 没有缓冲队列, 所以会造成死锁, 报如下错误fatal error: all goroutines are asleep - deadlock!
}

func size_max() {
	ch := make(chan int, 10)
	for i := 0 ; i < 2; i ++ {
		ch<-i
	}
	fmt.Println(len(ch), cap(ch))
}

func channel()  {
	size := 20
	ch := make(intChan, size) // 相当于chan的数组, 能放size个变量, 即当ch空时, 可进行size次ch<-i
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < size; i++ {
		ch <- i
	}
	fmt.Println(len(ch), cap(ch))
	go func() {
		for {
			start := time.Now()
			select {
			case v := <-ch:
				println("ch recv = ", v)
			// goroutine有时候会进入阻塞情况, 通过select设置超时处理可避免由于channel阻塞导致整个程序阻塞的发生
			case <-time.After(3 * time.Second):
				fmt.Printf("waittime = %s, break\n", time.Now().Sub(start).String())
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}

func channel_channel()  {
	chs := make(chan chan int, 10)
	ch := make(chan int, 1)
	chs <- ch
	fmt.Println(cap(chs), len(chs), chs)
	//go func() {
	//	ch <- 1
	//}()
	ch <- 1
	time.Sleep(1 * time.Second)
	fmt.Println(cap(chs), len(chs), chs)
}

func main()  {
	// size_0()
	size_max()
	// channel()
	// channel_channel()
}