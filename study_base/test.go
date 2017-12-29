package main

import (
	"fmt"
	importTest "goStudy/import_test"
)

func test1()  {
	importTest.PrintConfig()
}

type mystruct struct {
	a int
	b string
}

func main() {
	a := 1
	fmt.Println(&a)
	a,b := 1,2
	fmt.Println(&a, &b)
}
