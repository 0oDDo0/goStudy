package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 error 接口
func (de DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

// panic, 更加详细的报告异常, 包括函数调用顺序
func panic1() {
	panic2()
}

func panic2() {
	panic("panic test")
}

// 正常情况下panic异常, 只会报告异常位置和错误类型
func panic3() {
	panic4()
}
func panic4() {
	3 / 0
}

func main() {
	if false {
		panic3()
		panic1()
	}
	fmt.Println(2222)
	if true {
		// 正常情况
		if result, errorMsg := Divide(100, 10); errorMsg == "" {
			fmt.Println("100/10 = ", result)
		}
		// 当被除数为零的时候会返回错误信息
		if _, errorMsg := Divide(100, 0); errorMsg != "" {
			fmt.Println("errorMsg is: ", errorMsg)
		}
	}
}
