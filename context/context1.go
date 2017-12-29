package main

import (
	_ "context"
	"time"
	"context"
	"fmt"
)


/*

 Package context defines the Context type, which carries deadlines,
 cancelation signals, and other request-scoped(范围,视野,眼界) values across API boundaries
 and between processes.

 Incoming requests(请求) to a server should create a Context, and outgoing
 calls to servers should accept a Context. The chain of function
 calls between them must propagate(传送,传播) the Context, optionally replacing
 it with a derived(派生的) Context created using WithCancel, WithDeadline,
 WithTimeout, or WithValue. When a Context is canceled, all
 Contexts derived from it are also canceled.

 The WithCancel, WithDeadline, and WithTimeout functions take a
 Context (the parent) and return a derived Context (the child) and a
 CancelFunc. Calling the CancelFunc cancels the child and its
 children, removes the parent's reference to the child, and stops
 any associated timers(计时器). Failing to call the CancelFunc leaks(泄露) the
 child and its children until the parent is canceled or the timer
 fires. The go vet tool checks that CancelFuncs are used on all
 control-flow paths.

使用context时,注意的规则
1) 不要放在结构体中, Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it;
2) 函数首个参数, 命名为ctx, The Context should be the first parameter, typically named ctx;
3) 不要传nil作为context的父类, Do not pass a nil Context, even if a function permits it.
   Pass context. TODO,if you are unsure about which Context to use.
4) 线程安全, The same Context may be passed to functions running in different goroutines;
   Contexts are safe for simultaneous(同时) use by multiple goroutines.

context相当于线程监控器, 可优雅的结束线程

context.TODO()
    TODO返回一个非空, 空的上下文, 在目前还不清楚要使用的上下文或尚不可用时
context.Background(), root Context
    Background返回一个非空, 空的上下文. 这没有取消, 没有值, 并且没有期限, 它通常用在主函数, 初始化和测试,
    并作为输入的顶层上下文

Done方法在Context被取消或超时时返回一个close的channel,close的channel可以作为广播通知, 告诉给context相关的函数要停止当前工作
然后返回. 当一个父operation启动一个goroutine用于子operation,这些子operation不能够取消父operation.
Context可以安全的被多个goroutine使用. 开发者可以把一个Context传递给任意个goroutine, 然后cancel context时就能够通知到所有的goroutine

Err方法返回context为什么被取消
Deadline返回context何时会超时

WithCancel -> Context, func cancel
	Context:一个继承的Context,这个Context在父Context的Done被关闭时关闭自己的Done通道, 或者在自己被Cancel的时候关闭自己的Done
	func cancel: 一个取消函数cancel, 这个cancel用于取消当前的Context, 子协程存在map中, 循环kill

注意事项:
	不要将上下文存储在结构类型中; 而是将一个Context明确地传递给每个需要它的函数
	不要传递一个零上下文, 如果不确定要使用哪个上下文, 传递context.TODO。
*/

// 主动取消线程
func cancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx1, cancel1 := context.WithCancel(context.Background())
	// ctx ctx1是两个不同的子节点, 不会相互影响
	fmt.Println(&ctx, &ctx1)
	// 多个线程使用同一个Context时, Context.Done会发出多个信号, 让所有使用ctx的线程都会接收到停止的信号
	go doStuff(ctx, 1)
	go doStuff(ctx, 2)
	go doStuff(ctx, 3)
	go doStuff(ctx, 4)

	//10秒后取消doStuff
	time.Sleep(6 * time.Second)
	cancel1()
	time.Sleep(6 * time.Second)
	cancel()
	// 等待打印结果
	time.Sleep(5 * time.Second)
	fmt.Println(ctx.Err(), ctx.Value(0))
	fmt.Println(ctx.Deadline())
}

// 设置截止时间
func deadlineContext()  {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	go doStuff(ctx, 2)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

// 限定超时时间
func timeoutContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	go doStuff(ctx, 2)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context, second int64) {
	var times int64 = 0
	for {
		time.Sleep(time.Duration(second) * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(second, "done")
			return
		default:
			times++
			fmt.Println(second, "work", times)
		}
	}
}

func main() {
	cancelContext()
	// timeoutContext()
}
