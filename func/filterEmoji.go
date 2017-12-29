package main

import (
	"bytes"
	"unicode/utf8"
)

/* 常用的功能函数 */
func FilterEmoji(str string) string {
	if str == "" {
		return ""
	}
	var new_str bytes.Buffer
	new_str.Grow(len(str))
	for _, value := range str {
		if _, size := utf8.DecodeRuneInString(string(value)); size <= 3 {
			new_str.WriteRune(value)
		}
	}
	return new_str.String()
}

func main() {

}
