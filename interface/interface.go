package main

import (
	"fmt"
	"strconv"
)

/*
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
interface是一组抽象方法的集合,它必须由其他非interface类型实现, 而不能自我实现, 类似于c++中的虚基类
注:
	必须实现interface中所有的方法, interface才算作基类
	方法名, 参数, 返回值必须全保持一致
	interface的实例只能访问interface中的方法, 不能访问其派生类中的方法
空interface
	所有的类型都实现了空interface, 它可以存储任意类型的数值
	一个函数把interface{}作为参数时, 那么可以接受任意类型的值作为参数,
	如果一个函数返回interface{},那么也就可以返回任意类型的值
interface类似于c++中的void*, 可实现同一参数, 传递不同类型的值
*/
type Human struct {
	name string
	age  int
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human   //匿名字段
	company string
}

// Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, age = %d\n", h.name, h.age)
}

// Human实现Sing方法
func (h Human) Sing(song string) {
	fmt.Println("La la la la...", song)
}
func (h Human) eat() {
	fmt.Printf("%s is eating", h.name)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s\n", e.name, e.company)
}

// Interface Men被Human,Student和Employee实现 , 因为这三个类型都实现了这两个方法
// 必须实现interface中所有的方法, 只实现其中一个不能算作其实现
type Men interface {
	SayHi()
	Sing(song string)
}

func testInterface() {
	student := Student{Human{"david", 15}, "6-0"}
	employee := Employee{Human{"Tom", 25}, "google"}

	//定义Men类型的变量i
	var men Men
	// men能存储Student
	men = student
	fmt.Println("This is student, a Student:")
	men.SayHi()
	men.Sing("school bad")

	// men也能存储Employee
	men = employee
	fmt.Println("This is employee, an Employee:")
	men.SayHi()
	men.Sing("work hard")

	// 定义slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 2)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1] = student, employee
	for index, value := range x {
		// 只能调用相应的接口, 不能使用student的属性
		value.SayHi()
		value.Sing(fmt.Sprintf("%d", index))
		// value.speak() // 不能访问speak方法
	}
}

/*
空interface(interface{})不包含任何的method, 正因为如此, 所有的类型都实现了空interface.
空interface对于描述起不到任何的作用(因为它不包含任何的method),但是空interface在我们需要
存储任意类型的数值的时候相当有用, 因为它可以存储任意类型的数值,它有点类似于C语言的void*类型.

一个函数把interface{}作为参数时, 那么可以接受任意类型的值作为参数,
如果一个函数返回interface{},那么也就可以返回任意类型的值
*/
type nullInterface interface{}

func testNullInterface() {
	// a, b都为空interface
	var a nullInterface
	var b interface{}
	var i int = 5
	s := "Hello world"
	// 空interface可以存储任意类型的数值
	a = i
	b = s
	fmt.Println(a, b)
}

/*
interface函数参数, fmt.Println就是利用这种实现方式实现的, 利用这个特性, 可以让函数中的参数接受不同类型的值
任何实现了String方法的类型都能作为参数被fmt.Println调用
*/
func (h Human) String() string {
	return "{" + h.name + " - " + strconv.Itoa(h.age) + " age = " + fmt.Sprintf("%d", h.age) + "}"
}

func testArgsInterface() {
	h := Human{"ljs", 20}
	fmt.Println(h)
}

/*
Interface也应用于匿名字段, 即继承
*/
type Women interface {
	Men
}

func testWomen() {
	var women Women
	women = Student{Human{"ljs", 20}, "6-0"}
	women.SayHi()
	women.Sing("hello")
}

func main() {
	// testInterface()
	// testNullInterface()
	// testArgsInterface()
	testWomen()
}
