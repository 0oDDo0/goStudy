package main

import (
	"fmt"
	"reflect"
)

/*
If f is variadic(可变) with final parameter type ...T, then within the function the argument is equivalent to
a parameter of type[]T. At each call off, the argument passed to the final parameter is a new slice of type []T
whose successive elements are the actual(真实, 实际的) arguments, which all must be assignable(可分配的) to the type T.
The length of the slice is the number of arguments bound(范围,限制) to the final parameter and may differ for each call site.
*/

// 可变参数, 类型为切片
func print(args ...interface{}) {
	fmt.Println(args, reflect.TypeOf(args))
	for _, i := range args {
		fmt.Print(i,",")
	}
	fmt.Println()
	// 如果放在一起, 会报错, 参数过多
	fmt.Print("args="); fmt.Println(args...)// ...相当于遍历切片了,即fmt.Println(args[0], args[1], ..., args[i])
}

func runPrint()  {
	print(1);
	print(1, []int{2,3});
	print(1, 2, 3);
}

func slice(args []interface{})  {
	fmt.Println(args, reflect.TypeOf(args))
	for _, i := range args {
		fmt.Print(i,",")
	}
	fmt.Println()
	// 如果放在一起, 会报错, 参数过多
	fmt.Print("args="); fmt.Println(args...)// ...相当于遍历切片了,即fmt.Println(args[0], args[1], ..., args[i])
}



func main() {
	// runPrint()
	slice([]interface{}{1, 2, 3})
	fmt.Println()
}
