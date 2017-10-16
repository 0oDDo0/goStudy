package main

/*
for三种写法
1) for init; condition; post { }, 等价于c++中的for(int i = 0; i < 10; i++)
2) for condition { }, 等价于c++中的for(;i < 10;)
3) for { }, 死循环, 等价于c++中的for(;;)
*/
func test_for() {
	i := 0
	for i = 0; i < 2; i++ {
		println("i1 =", i)
	}
	println()
	for i < 4 { // 等价于 for ; i < 4;
		println("i2 =", i)
		i++
	}
	println()
	numbers := [6]int{1, 2, 3, 5}
	// 类似于Python中的enumerate(numbers)
	// for循环的 range 格式可以对 slice,map,数组,字符串等进行迭代循环
	for index, value := range numbers {
		println("第", index, "位 x 的值 =", value)
	}
	println()
	// 遍历可迭代的对象, _的使用与Python类似
	for _, value := range numbers {
		println("value =", value)
	}
}
func main() {
	test_for()
}
