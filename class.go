package main

import (
	"fmt"
	"math"
)

/*
method, 函数的另一种形态, 带有接收者的函数
method附属在一个给定的类型上的,他的语法和函数的声明语法几乎一样,只是在func后面增加了一个receiver(也就是method所依从的主体)
func (r ReceiverType) funcName(parameters) (results)
注:
  虽然method的名字一模一样,但是如果接收者不一样, 那么method就不一样
  method里面可以访问接收者的字段
  调用method通过.访问, 就像struct里面访问字段一样
  Receiver是以值传递,而非引用传递
  Receiver可以为指针, https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md
	1) 如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method,
	   而不需要&V去调用这个method;
	2) 如果一个method的receiver是T,你可以在一个*T类型的变量P上面调用这个method,
	   而不需要*P去调用这个method;
	so不用担心调用的指针的method还是不是指针的method
  私有公有也是通过首字母的大小写来决定
*/

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
// 独立的函数
func area(r *Rectangle) float64 {
	return r.width * r.height
}
// method,  area()是属于Rectangle的方法
// Rectangle存在以下属性:height,width,area()
func (r Rectangle) area() (float64){
	return r.height * r.width
}
// 虽然名字一样, 但是接收者不一样, 所以是不一样的method
func (c Circle) area() (float64){
	return math.Pi * c.radius * c.radius
}
// c *Circle和c Circle,可以认为是同一类型的接收者,编译器会自动将c Circle转换为对应的指针
// 所以method不可以同名
// c Circle 为值传递, 不会影响到原始对象
// c *Circle为引用传递, 会影响到原来的对象
func (c *Circle) area_ptr() (float64){
	c.radius = 3
	return math.Pi * c.radius * c.radius
}
type String string
// 不可以使用string, 报错
func (i String) area() (String){
	return i + " area"
}

func main() {
	r := Rectangle{12, 2}
	c := Circle{2}
	s := String("a")
	fmt.Println("Area of r is: ", area(&r))
	fmt.Println("Area of r is: ", r.area())
	fmt.Println("Area of c is: ", c.area())
	fmt.Println("Area of c is: ", c.area_ptr())
	fmt.Println("Area of c is: ", c.area())
	fmt.Println("Area of s is: ", s.area())
}