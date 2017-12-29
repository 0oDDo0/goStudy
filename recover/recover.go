package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func generatePanic() {
	panic("")
}

// 将捕获的异常作为函数返回值
func getError1() (err error) {
	// If there is a panic we need to recover in a deferred func
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recover error")
		}
	}()

	generatePanic()
	fmt.Println("run here") // 发生异常后,下面的语句不会再执行, 所以不会打印这句
	return
}

func main()  {
	fmt.Println(getError1())
}