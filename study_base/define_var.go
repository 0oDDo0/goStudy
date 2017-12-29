/*
第一行代码 package main 定义了包名, 必须在源文件中<非注释>的第一行指明文件属于哪个包
go语言中的  <变量类型> 放在后面
*/
package main

// 导入fmt包, 如果导入了包, 但是未使用, 编译会出现错误
import (
	"fmt"
)

/*
当标识符(包括常量,变量,类型,函数名,结构字段等)以大写字母开头时, 这种形式的标识符对象就可以被外部包的代码所使用, 类似于c++中的public
标识符如果以小写字母开头, 则对外是不可见的, 类似于c++中的protected
*/

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	author string = "lis"
	age    int    = 25
)

// 定义全局常量, 类型可以不写, 编译器会自己推导
const target, order = "study go", 1

/*
与包名同名的函数是程序开始执行时需要执行的函数,如果有init()函数则先执行init()函数
*/
func init() {
	fmt.Println("init, hello")
}
func define_var() {
	// Go语言中变量的声明必须使用空格隔开, 对于已声明但未使用的变量会在编译阶段报错
	// 变量声明的方式
	var age1 int = 1
	// 编译器自行推导变量类型
	var age2 = 2
	// 先指定变量类型, 然后赋值, 如果不赋值, 使用默认值
	var age3 int
	age3 = 3
	// 初始化声明, 使用:=, 省略var, 左侧的变量不应该是一句声明过的, "这种不声明格式的变量只能在函数体中定义"
	age4 := 4 // 等价于var age4 = 2, 如果age4已经声明了,则会报错
	println("fun, hello age = ", age1, age2, age3, age4)
	// 多变量声明的格式
	var age5, age6 = 5, 6
	// _的含义与Python相同, 表示此变量未使用, 任何赋予它的值都会被丢弃
	_, age7, age8 := 0, 7, 8

	println("fun, hello age = ", age5, age6, age7, age8)
}

// 值和指针,和c++相似
func value_and_ptr() {
	var i = 1
	// var i = 1
	var ptr1 = &i
	*ptr1 = 2
	println(i, ptr1)
	var ptr2 *int = ptr1
	*ptr2 = 3
	println(i, ptr1)
}

// 使用简式声明重复声明变量, 你不能在一个单独的声明中重复声明一个变量,但在多变量声明中这是允许的,其中至少要有一个新的声明变量
// 重复变量需要在相同的代码块内,否则你将得到一个隐藏变量
func define_mul_var()  {
	{
		a := 1
		fmt.Println(&a)
		a,b := 2,3 // a还是原来的a, 地址相同
		fmt.Println(&a, &b)
	}
	{
		x := 1
		fmt.Println(x)     // print 1
		{
			fmt.Println(x) // print 1
			x := 2
			fmt.Println(x) // print 2
		}
		fmt.Println(x)     // print １
	}
}
func main() {
	define_mul_var()
	//define_var()
	//value_and_ptr()
	//println("main, hello, author =", author, age, "target =", target, "order =", order)
}
