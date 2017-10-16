package main

import "fmt"

/*
make用于内建类型(map,slice,channel)的内存分配, new用于各种类型的内存分配
make返回实例, make只能创建slice,map,channel, 并且返回一个有初始值(非零)的T类型
	本质来讲,导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化,
	例如,一个slice,是一个包含指向数据(内部array)的指针,长度和容量的三项描述符;
	在这些项目被初始化之前,slice为nil,对于slice来说,make初始化了内部的数据结构,填充适当的值
new返回指针, new(T)分配了零值填充的T类型的内存空间, 并且返回其地址, 即一个*T类型的值
*/

func correct_array1(a []int)  {
	a[0] = 1
}

func correct_array2(a *[8]int)  {
	a[0] = 1
}
func main() {
	a := make([]int, 8)
	b := new([8]int)
	fmt.Println(len(a), len(b), a, b)
	correct_array1(a)
	correct_array2(b)
	fmt.Println(len(a), len(b), a, b)
}
