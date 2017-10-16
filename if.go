package main

import "fmt"

func test_if() {
	/*
	go特性, 条件判断语句中, 允许声明一个变量, 作用域仅限在if中
	*/
	var a = 10
	if a > 0 && a < 20 {
		fmt.Println("condition true")
	} else { // else不能换行
		fmt.Println("condition false")
	}
	if x := 10; x == a {
		fmt.Println("x == a")
	}
}
func test_switch(x int) {
	var age int
	// case后面不需要加break, 每个case后面默认加了break,
	// 如果想继续执行后面case的代码, 可以添加fallthrough, 不管后面case条件是否满足, 都会执行
	// case后面必须为相同类型的变量
	switch x {
	case 1:
		age = 1
	case 2:
		age = 2
		// 当age==2时,还会继续执行后面的case, 后面case不满足条件也会执行, 即age=3
		fallthrough
	// 多个符合的值时, 用,分开
	case 3, 4, 5:
		age = 3
	default:
		age = 0
	}
	println("age =", age)
}

func test_select() {
	/*
		http://www.runoob.com/go/go-select-statement.html
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
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func main() {
	// test_if()
	test_switch(2)
	// test_select()
}
