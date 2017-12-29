package main

import "fmt"

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

/*
Go语言中"range"关键字用于for循环中迭代数组(array),切片(slice),通道(channel),集合(map)中的元素;
*/

func test_range() {
	//使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 需要使用index和value
	for index, value := range nums {
		if value == 3 {
			fmt.Println("index:", index)
		}
	}
	// range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("key = %s -> value = %s\n", key, value)
	}
	//range也可以用来枚举Unicode字符串, 第一个参数是字符的索引,第二个是字符(Unicode的值)本身。
	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}
func main() {
	test_for()
}
