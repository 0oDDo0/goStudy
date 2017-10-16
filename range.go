package main
import "fmt"
/*
Go语言中"range"关键字用于for循环中迭代数组(array),切片(slice),通道(channel),集合(map)中的元素;
*/
func main() {
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