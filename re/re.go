package main

import (
	"fmt"
	"regexp"
)

// 按照标点符号切分
func fun1() {
	str := "a,b.c...d!e.?f:g,"
	var re *regexp.Regexp
	re = regexp.MustCompile(`[?.!,:;]+`)
	strs := re.Split(str, -1)
	for index, val := range strs {
		fmt.Println(index, val)
	}
}

func fun2() {
	str := "ab. 11 &.11 "
	var re *regexp.Regexp
	re = regexp.MustCompile(`[a-z0-9_'.-]+`)
	strs := re.FindAllString(str, -1)
	for index, val := range strs {
		fmt.Println(index, val)
	}
	//fmt.Println(re.ReplaceAllString(str, " "))
}

func main() {
	fun2()
}
