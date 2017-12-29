package main

import "fmt"

// 常量, 也就是在程序编译阶段就确定下来的值,而程序在运行时无法改变该值. 在Go程序中,常量可定义为数值,布尔值或字符串等类型
// const identifier [type] = value
const author, age = "ljs", 25

// 枚举类型, 0,1,2分别代表未知性别,女性和男性
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	const pi = 3.14
	var r = 2.0
	println(int(pi * r * r))
	/*
		Go里面有一个关键字iota, 这个关键字用来声明enum的时候采用, 它默认开始值是0, const中每增加一行加1

		除非被显式设置为其它值或iota, 每个const分组的第一个常量被默认设置为它的0值,
		第二及后续的常量被默认设置为它前面那个常量的值, 如果前面那个常量的值是iota, 则它也被设置为iota。
	*/
	{
		const (
			x = iota // x == 0
			y = iota // y == 1
			z = iota // z == 2
			w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
		)

		const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

		const (
			h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
		)

		const (
			a       = iota //a=0
			b       = "B"
			c       = iota             //c=2
			d, e, f = iota, iota, iota //d=3,e=3,f=3
			g       = iota             //g = 4
		)
		fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	}
}
