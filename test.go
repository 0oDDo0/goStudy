package main

import (
	"fmt"
	"unicode/utf8"
	"bytes"
)

func FilterEmoji(str string) string {
	if str == ""{
		return ""
	}
	var new_content bytes.Buffer
	new_content.Grow(len(str))
	for _, value := range str {
		if _, size := utf8.DecodeRuneInString(string(value));size <= 3 {
			new_content.WriteRune(value)
		}
	}
	return new_content.String()
}

func main() {
	value := 2
	if value := 1; value > 0 {
		fmt.Println(value)
	}
	fmt.Println(value)
}