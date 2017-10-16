/* go中使用的技能 */
package main
// 同时声明多个常量,变量,或者导入多个包时，可采用分组的方式进行声明。
//import_test(
//	"fmt"
//	"os"
//)

const(
	i = 100
	pi = 3.1415
	prefix = "Go_"
)

var(
	i int
	pi float32
	prefix string
)

/*
Go之所以会那么简洁，是因为它有一些默认的行为:
	1 大写字母开头的变量是可导出的, 也就是其它包可以读取的, 是公有变量; 小写字母开头的就是不可导出的, 是私有变量;
	2 大写字母开头的函数也是一样, 相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
 */
func main() {
}
