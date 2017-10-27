package main

import (
	"github.com/deckarep/golang-set"
	"fmt"
)

var a = 1

func test()  {
	a = 2
}
func main()  {
	test()
	fmt.Println(a)
	kide := mapset.NewSet(1, 2)
	fmt.Println(kide)
	kide.Add("xiaorui.cc")
	kide.Add("blog.xiaorui.cc")
	kide.Add("vps.xiaorui.cc")
	kide.Add("linode.xiaorui.cc")
	kide.Add("Biology")
	for val := range kide.Iter() {
		fmt.Println("SET = ", val)
	}

	special := []interface{}{"Biology", "Chemistry"}
	scienceClasses := mapset.NewSetFromSlice(special)

	address := mapset.NewSet()
	address.Add("beijing")
	address.Add("nanjing")
	address.Add("shanghai")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	//一个并集的运算
	allClasses := kide.Union(scienceClasses).Union(address).Union(bonusClasses)

	//是否包含"Cookiing"
	fmt.Println(scienceClasses.Contains("Cooking")) //false

	//两个集合的差集
	fmt.Println(allClasses.Difference(scienceClasses)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	//两个集合的交集
	x := scienceClasses.Intersect(kide) //Set{Biology}
	fmt.Println("& = ", x, x.Cardinality())

	//有多少基数
	fmt.Println(bonusClasses.Cardinality()) //2
}

