package main

import "fmt"

type human struct {
	name string
}

func (h human)say() bool {
	fmt.Println("my name is", h.name)
	return false
}


func (h human)speak() bool {
	fmt.Println("speak, my name is", h.name)
	return false
}

// 测试只实现借口中的一个方法, 还算不算基类, 结果, 必须实现interface中的全部方法
type human_interface interface {
	say() bool
	// song()
}

func TestInterface()  {
	var a human_interface
	a = human{"ljs"}
	a.say()
}

func main() {
	TestInterface()
}
