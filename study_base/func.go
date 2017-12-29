package main

import "fmt"

/*
函数的定义, 默认情况下,Go语言使用的是值传递
func function_name( [parameter list] ) ([return_types list]) {
   函数体
}
parameter list = var1 type, var2 type
return_types list = ret1 type1, ret2 type2..., 可以返回多个参数
func使用
	关键字func用来声明一个函数func_name
	函数可以有一个或者多个参数, 每个参数后面带有类型, 通过","分隔
	参数类型相同时, 可省略, 如a int, b int -> a, b int, a的类型与离他最近的类型相同
	函数可以返回多个值
	返回值中可以不声明变量, 直接写返回的类型
	如果只有一个返回值且不声明返回值变量, 那么可以省略返回值上的()
	如果没有返回值, 直接省略最后的返回信息
	如果有返回值, 那么必须在函数的外层添加return语句
	不支持重载
*/

func max(num1 int, num2 int) int { // 等价于num1, num2 int
	/* 声明局部变量 */
	if num1 > num2 {
		return num1
	}
	return num2
}

// 返回多个值
func swap1(x string, y string) (string, string) {
	return y, x
}

// 最好命名返回值,因为不命名返回值,虽然使得代码更加简洁了,但是会造成生成的文档可读性差
// 命名了返回值时, 直接return 即可
func swap2(x string, y string) (m string, n string) {
	m, n = y, x
	return // 等价于return m, n
	// return m,n
}

// 支持变参, 即参数个数不定
func multi_pra(args ...string) {
	fmt.Println(args) // args = string的slice
}

func main() {
	fmt.Println(wrapper()(1)) // 先返回一个函数, 然后在调用函数
	if false {
		fmt.Printf("最大值是 : %d\n", max(1, 2))
		// 函数也可以当做变量来传递
		var a, b string
		Swap := swap1
		a, b = Swap("a", "b")
		a, b = swap2(a, b)
		fmt.Println(a, b)
		multi_pra(a, b)

		slice := []int{1, 2, 3, 4, 5, 7}
		fmt.Println("slice = ", slice)
		odd := filter(slice, isOdd) // 函数当做值来传递了
		fmt.Println("Odd elements of slice are: ", odd)
		even := filter(slice, isEven) // 函数当做值来传递了
		fmt.Println("Even elements of slice are: ", even)
	}
}

type testInt func(int) bool // 声明了一个函数类型

// 奇数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

// 偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f testInt) (result []int) {
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// 闭包
func wrapper() func(i int) int {
	return func(i int) int {
		return i + 1
	}
}
