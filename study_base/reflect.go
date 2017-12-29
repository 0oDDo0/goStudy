package main

import (
	"fmt"
	"reflect"
)

/*
反射就是能检查程序在运行时的状态, 获取变量类型等信息
反射一个类型的值(这些值都实现了空interface), 需要转化成reflect对象(reflect.Type or reflect.Value, 不同情况调用不同	函数)
  1) t := reflect.TypeOf(i)
    得到类型的元数据,通过t我们能获取类型定义里面的所有元素
  2) v := reflect.ValueOf(i)
    得到实际的值,通过v我们获取存储在里面的值,还可以去改变值
*/

func testReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
func main() {
	testReflect()
}
