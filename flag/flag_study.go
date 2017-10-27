package package_study

import (
	"flag"
	"fmt"
)

// main.go调用此函数, 执行命令: go run main.go -x 1 -z 2 -y 1, 将-z放在-y后面执行不成功, why?
func TestFlag() {
	var var_z int
	var_x := flag.Int("x", 0, "value of x") // type(x) = *int, 注册参数x, 如果命令行中有此参数, 赋给var_x
	var_y := flag.Bool("y", false, "value of y")
	flag.IntVar(&var_z, "z", 0, "value of z")
	flag.Parse()
	fmt.Println(*var_x, *var_y, var_z)
}