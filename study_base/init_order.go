package main

import "fmt"

/*
初始化顺序, import -> const -> var -> init() ->main
*/
var b = a

const a = 1

func init() {
	fmt.Println("a =", a, "b =", b)
	b = 2
}

func main() {
	fmt.Println("a =", a, "b =", b)
}
