package import_test

import "fmt"

/*
同一个包中, 所有的文件可以视为由一个大的文件拆分而成.
所以同一个包中的文件可以直接使用另一个文件中的变量和属性, 也不能重复定义变量和方法等.
同一个包中可以写多个init, 但是最好不要这么干... 执行顺序不确定
*/
var (
	a = 0
	b = 1
	A = 0
	B = 1
)

const (
	c = 2
	d = 3
	C = 2
	D = 3
)

func init() {
	fmt.Println("config init")
	a = 1
}

func PrintConfig() {
	fmt.Println(fmt.Sprintf("config.go, a = %d b = %d c = %d d = %d", a, b, c, d))
}
