package main

import "fmt"
/*
struct是可以通过匿名字段继承, method也是可以继承的
*/
type Human struct {
	name string
	age int
}

type Student struct {
	Human // 匿名字段
	grade string
}


func (h *Human) sayHi() {
	fmt.Printf("Hi, I am %s, age = %s\n", h.name, h.age)
}

func (h *Human) overWrite() {
	fmt.Printf("human over write\n")
}
// 重写overWrite方法
func (h *Student) overWrite() {
	fmt.Printf("student over write\n")
}

func main() {
	student := Student{Human{"ljs", 20}, "6-0"}
	student.sayHi()
	student.overWrite()
}
