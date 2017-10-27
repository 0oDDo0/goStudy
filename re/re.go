package main

import (
	"regexp"
	"fmt"
)

func fun1() {
	str := "a,b.c...d!e.?f:g,"
	var re *regexp.Regexp
	re = regexp.MustCompile(`[?.!,:;]+`)
	strs := re.Split(str, -1)
	for index, val := range strs{
		fmt.Println(index, val)
	}
}

func fun2() {
	str := "ab Aa-bVVV a_bCCC a'b"
	var re *regexp.Regexp
	re = regexp.MustCompile(`[^a-z0-9_'-]+`)
	strs := re.FindAllString(str, -1)
	for index, val := range strs{
		fmt.Println(index, val)
	}
	fmt.Println(re.ReplaceAllString(str, " "))
}

func main()  {
	fun2()
}
