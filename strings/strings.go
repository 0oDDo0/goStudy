package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	str := "a b  c	d\ne		f"
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc(str, unicode.IsSpace))
	split := strings.Split(str, "\t")
	for index, val := range split {
		fmt.Println(index, val)
	}
	fmt.Println(strings.Join(split, " "))
	fmt.Println(strings.TrimSpace("sdada \n adad"))
	fmt.Println(strings.TrimSpace("sdada \n"))

}