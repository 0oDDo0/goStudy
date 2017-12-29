package main

import (
	"encoding/json" // 利用反射解析json
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"reflect"
	//"time"
)

/*
json 在生成序列和反序列化的时候, 都用到了反射, 所以效率会较差
提高效率方法: 如调用reflect, 提前为所有的struct生成好序列化代码, 这样只需要在生成的时候使用reflect,而运行时序列化
通过使用生成的代码, 效率得到大幅提高
// https://github.com/pquerna/ffjson
ffjson提前为Go中的结构生成静态MarshalJSON和UnmarshalJSON函数. 生成的函数减少了反射的操作, 进而减少序列化,反序列化的时间
因为提前生成了对应结构的序列,反序列化函数, 从而减少了反射的次数
化
*/
func test_json(x interface{}) {
	pack, _ := json.Marshal(x) // obj -> json
	packStr := string(pack)
	var unpack interface{}
	json.Unmarshal(pack, &unpack) // json -> obj
	fmt.Println("type(x) = ", reflect.TypeOf(x), "x = ", x)
	fmt.Println("type(packStr) = ", reflect.TypeOf(packStr), "result = ", packStr)
	fmt.Println("type(unpack) = ", reflect.TypeOf(unpack), "result = ", unpack, "\n")
}

func test_ffjson(x interface{}) {
	pack, _ := ffjson.Marshal(x) // obj -> json
	packStr := string(pack)
	var unpack interface{}
	ffjson.Unmarshal(pack, &unpack) // json -> obj
	fmt.Println("type(x) = ", reflect.TypeOf(x), "x = ", x)
	fmt.Println("type(packStr) = ", reflect.TypeOf(packStr), "result = ", packStr)
	fmt.Println("type(unpack) = ", reflect.TypeOf(unpack), "result = ", unpack)
	fmt.Println()
}

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"response2_page"`
	Fruits []string `json:"response2_fruits,omitempty"`
}

func test_json_struct() {
	// 结构体 -> json
	res1D := &Response1{Page: 1}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{Page: 1}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	res2D.Fruits = []string{"apple", "peach", "pear"}
	res2B, _ = json.Marshal(res2D)
	fmt.Println(string(res2B))

	// json -> 结构体
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response1{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("res", res, res.Page, res.Fruits)
}

func main() {
	res1D := &Response1{Page: 1, Fruits:[]string{"a", "b"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// m := map[string][]string{"a": {"1", "2"}, "b": {"3", "4"}}
	// test_json_struct()
	//// var objects = []interface{}{1, true, 1.111, "abcd", []int{1,2,3}, map[string]interface{}{"a":1, "b":2}}
	//var objects = []interface{}{map[string]interface{}{"a":1, "b":2}}
	//// start := time.Now()
	//for i := 0; i < 1; i++ {
	//	for _, val := range objects {
	//		test_json(val)
	//	}
	//}
	//end := time.Now()
	//fmt.Println("run time", end.Sub(start)) // 98ms
	//start = time.Now()
	//for i := 0; i < 1000; i++ {
	//	for _, val := range objects {
	//		test_ffjson(val)
	//	}
	//}
	//end = time.Now()
	//fmt.Println("run time", end.Sub(start)) // 74ms, ffjson
}
