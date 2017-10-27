package main

import "fmt"

/* interface{} -> 指定的类型, 类型断言
https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
https://studygolang.com/articles/3314
推荐使用以下做转换
if str, ok := data.(string); ok {
	do something
} else {
	exception handle
}
*/

func translate() []string {
	var s interface{} = []string{"1", "2"}
	// 如果为指定的类型, flag = true, trans为对应类型的值; else, flag = false
	trans, flag := s.([]string)
	if flag{
		return trans
	}
	return nil
}

// 虽然可以直接用s.([]string)来转换interface{},但是断言失败会导致panic error
func translate_right() []string {
	var s interface{} = []string{"1", "2"}
	return s.([]string)
}
// panic: interface conversion: interface {} is int, not []string
func translate_error() []string {
	var s interface{} = 1
	return s.([]string)
}

func main () {
	fmt.Println(translate())
	fmt.Println(translate_right())
	fmt.Println(translate_error())
}