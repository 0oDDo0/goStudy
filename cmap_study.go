package main

import (
	"github.com/orcaman/concurrent-map"
	"fmt"
)

/*
线程安全map(加了锁的map)
将存储区域分片, 每片使用一把锁, 提升效率,
*/

func study()  {
	var dict cmap.ConcurrentMap = cmap.New()
	tmp := make(map[string]int)
	tmp["key1"] = 1
	tmp["key2"] = 2
	dict.Set("key1", "value1")
	dict.Set("key2", 2)
	dict.Set("key3", tmp)
	fmt.Println(dict.Get("key3"))
	fmt.Println(dict.GetShard("key3"))
}

func main()  {
	study()
}