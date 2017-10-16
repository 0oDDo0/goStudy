package main

import (
	"fmt"
	"runtime"
	"time"
)
/*
并发编程
goroutine是Go并行设计的核心.goroutine说到底其实就是协程, 但是它比线程更小, 十几个goroutine可能体现在底层就是五六个线程,
Go语言内部帮你实现了这些goroutine之间的内存共享. 执行goroutine只需极少的栈内存(大概是4~5KB),正因为如此,
可同时运行成千上万个并发任务. goroutine比thread更易用,更高效,更轻便.
*/
/*
runtime包中有几个处理goroutine的函数:
	1) Goexit, 退出当前执行的goroutine, 但是defer函数还会继续调用;
	2) Gosched, 让出当前goroutine执行权限, 调度器安排其他等待任务运行, 并在下次某个时候从该位置恢复执行;
	3) NumCPU, 返回CPU核数量;
	4) NumGoroutine, 返回正在执行和排队的任务总数;
	5) GOMAXPROCS, 用来设置可以并行计算的CPU核数的最大值, 并返回之前的值;
 */
func testParallel(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(fmt.Sprintf("%s %d", s, i))
	}
}
/*
goroutine(协程)运行在相同的地址空间, 因此访问共享内存必须做好同步.
Go提供了一个很好的通信机制channel, 可以通过它发送或者接收值, 这些值只能是特定的类型:channel类型.
定义一个channel时, 也需要定义发送到channel的值的类型, 必须使用make创建channel

channel, 通过操作符<-来接收和发送数据
ch <- v    发送v到channel ch
v := <-ch  从ch中接收数据, 并赋值给v

channel, 非缓存类型, 即
默认情况下, channel接收和发送数据都是阻塞的, 除非另一端已经准备好, 这样就使得Goroutines同步变的更加的简单,而不需要显式的lock.
所谓阻塞, 也就是如果读取(value := <-ch)它将会被阻塞, 直到有数据接收;
其次, 任何发送(ch<-5)将会被阻塞直到数据被读出. 无缓冲channel是在多个goroutine之间同步很棒的工具
 */
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println("calumniating a = ", a)
	c <- total  // send total to c
	fmt.Println("return a = ", a)

}

func testChannels()  {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	fmt.Println("waiting")
	x, y := <-c, <-c  // receive from c
	fmt.Println("cal finish")
	// 等待打印
	time.Sleep(10)
	fmt.Println(x, y, x + y)
}

/*
缓存channel, Go也允许指定channel的缓冲大小,就是channel可以存储多少元素. 如
ch:= make(chan bool, 4),创建了可以存储4个元素的bool型channel,前4个元素可以无阻塞的写入.
当写入第5个元素时, 代码将会阻塞, 直到其他goroutine从channel中读取一些元素,腾出空间

ch := make(chan type, value)
	1) 当value=0时,channel是无缓冲阻塞读写的;
	2) 当value>0时,channel有缓冲,是非阻塞的,直到写满value个元素才阻塞写入;
*/

func testBufferChannel()  {
	c := make(chan int, 2)// 修改2为1就报错, 因为进程会一直阻塞, >=2即可
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/*
range和close
*/
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
/*
for i := range c能够不断的读取channel里面的数据,直到该channel被显式的关闭. 生产者通过内置函数close显示关闭channel,
关闭channel之后就无法再发送任何数据了;
在消费方可以通过语法v, ok := <-ch测试channel是否被关闭, 如果ok返回false,那么说明channel已经没有任何数据并且已经被关闭.
ps:
	1) 应该在生产者的地方关闭channel, 而不是消费的地方去关闭它, 这样容易引起panic;
	2) channel不像文件之类, 不需要经常去关闭, 只有当你确实没有任何发送数据了,或者想显示关闭时才close channel
 */
func testRangeClose()  {
	c := make(chan int, 2)
	go fibonacci(10, c)
	for i := range c {
		fmt.Print(i, " ")
	}
}

func selectFibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
/*
如果存在多个channel时, 通过select可以监听多个channel上的数据流动;
1) select默认是阻塞的, 只有当监听的channel中有发送或接收可以进行时才会运行;
2) 当多个channel都准备好的时候, select是随机的选择一个执行的;
3) 在select有default语法, default就是当监听的channel都没有准备好时, 默认执行的代码, 此时select不再阻塞等待channel

以下描述了 select 语句的语法:
	1) 每个case都必须是一个通信
	2) 所有channel表达式都会被求值
	3) 所有被发送的表达式都会被求值
	4) 如果任意某个通信可以进行,它就执行;其他被忽略
	5) 如果有多个case都可以运行,Select会随机公平地选出一个执行.其他不会执行。
	   否则:
	       如果有default子句, 则执行该语句.
		   如果没有default字句,select将阻塞,直到某个通信可以运行;Go不会重新对channel或值进行求值
 */
func testSelect()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c, " ")
		}
		quit <- 0
	}()
	selectFibonacci(c, quit)
}

/*
超时
*/
func testTimeOut()  {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c:
				println("c receive = ", v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	c <- 10 // 发送数据
	<- o    // 接收数据, 阻塞当前进程, 等待超时
}

func main() {
	// go关键字很方便的就实现了并发编程
	if false {
		go testParallel("world")
		// 开一个新的Goroutines执行
		testParallel("hello") // 当前Goroutines执行
	}
	// testChannels()
	// testBufferChannel()
	// testRangeClose()
	// testSelect()
	testTimeOut()
}