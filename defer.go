package main

import "fmt"

/*
延迟(defer)语句,可在函数中添加多个.函数执行到最后时,defer语句会按照逆序执行,最后函数返回.
*/

func func_defer()  {
	for i := 0; i < 5; i++ {
		// defer语句按照栈的顺序执行, 即先定义的后执行
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("before return")
	// defer会在return之前执行
}
func main() {
	func_defer()
}
