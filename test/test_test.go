package test

import (
	"testing"
	"fmt"
)

/*
性能测试, 测试案例
*/
/*
  1.单元测文件名为*_test.go, 所以测试代码和包中的业务代码是分开的
    *_test.go程序不会被普通的Go编译器编译, gotest会编译普通程序和测试程序
  2.测试函数命名:Test*(t *testing.T), *的第一个字母大写
    T是传给测试函数的结构类型,用来管理测试状态,支持格式化测试日志

interface:
	1) func (t *T) Fail()
	   标记测试函数为失败,然后继续执行剩下的测试案例
	2）func (t *T) FailNow()
	   标记测试函数为失败并中止执行, 文件中别的测试也被略过, 继续执行下一个文件
	3) func (t *T) Log(args ...interface{})
    	args 被用默认的格式格式化并打印到错误日志中
	4) func (t *T) Fatal(args ...interface{})
       3) + 2), 先执行3), 然后执行2)
*/

// 使用方法
func TestA(t *testing.T) {
	fmt.Println("Test")
	if 1 < 2 {
		t.Fatal("test log1")
	}
	if 1 < 2 {
		t.Log("test log1")
		t.Fail()
	}
}

func fun(s string) string{
	return s
}
func TestB(t *testing.T) {
	type tests struct {
		right string // 正确结果
		input string // 输入
	}
	compares := []tests{{"1", "2"}, {"2", "2"}}
	for _, compare := range compares {
		ret := fun(compare.input)
		if ret != compare.right {
			t.Errorf("input = %s func result = %s right = %s", compare.input, ret, compare.right)
		}
	}
}

// 测试一个函数的平均执行时间
func BenchmarkC(b *testing.B) {
	b.SetParallelism(10000) // 也可以不设置并行数量, Go会选择最优值
	for i := 1; i < 10000;  {
		i++
	}
	/*
	result BenchmarkC    10000000    282 ns/op
	意味着循环执行了 10000000 次, 每次循环花费 282 纳秒(ns)。
	 */
}