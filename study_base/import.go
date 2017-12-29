package main

/*
import_test "fmt", 去GOROOT环境变量指定目录下去加载该模块
import_test "./model", 相对路径导入,即当前文件同一目录的model目录, 但是不建议使用这种方式
import_test "shorturl/model", 绝对导入, 加载gopath/src/shorturl/model模块
*/
/*
import的其他使用技巧
1) 点操作
	import_test . "fmt", 相当于Python中的from fmt import_test *, 即使用fmt中的函数时, 不用加包名
2) 别名操作
	import_test another_name "fmt", 相当于fmt的别名
3) _操作
	_操作虽然引入该包, 但是不直接使用包里面的函数, 而是调用了该包里面的init函数

import流程
包的导入过程说明
程序的初始化和执行都起始于main包, 如果main包还导入了其它的包, 那么就会在编译时将它们依次导入.
有时一个包会被多个包同时导入,那么它只会被导入一次(例如很多包可能都会用到fmt包, 但它只会被导入一次, 因为没有必要导入多次).
当一个包被导入时, 如果该包还导入了其它的包, 那么会先将其它包导入进来, 然后再对这些包中的包级常量和变量进行初始化,
接着执行init函数(如果有的话), 依次类推. 等所有被导入的包都加载完毕了, 就会开始对main包中的包级常量和变量进行初始化,
然后执行main包中的init函数(如果存在的话), 最后执行main函数

golang的package的特点:
  1) go的package不局限于一个文件, 可以由多个文件组成; 组成一个package的多个文件,编译后实际上和一个文件类似;
     组成包的不同文件相互之间可以直接引用变量和函数，不论是否导出;
     因此,组成包的多个文件中不能有相同的全局变量和函数,因为一个目录下的多个文件可以视为一个文件, 只不过格式上进行了细分
  2) go不要求package的名称和所在目录名相同, 但是最好保持相同, 否则容易引起歧义.
     因为引入包的时候, go会使用子目录名作为包的路径, 而你在代码中真正使用时, 却要使用你package的名称
     <导入的是路径名, 真正使用其中的属性时, 需要使用包名>
  3) 每个子目录中只能存在一个package, 否则编译时会报错, 即一个目录下, package的名字要一致
  4) go的package是以绝对路径GOPATH来寻址的, 不要用相对路径来import
*/

import (
	ip "./import_test"
	"fmt"
)

func init() {
	fmt.Println("main init")
}
func main() {
	// 大写开头的属性外部才能使用
	fmt.Println(fmt.Sprintf("config.go, A = %d B = %d C = %d D = %d", ip.A, ip.B, ip.C, ip.D))
	ip.PrintConfig()
	ip.PrintUse()
	fmt.Println("main")
}

/*
config init, 因为config.go中定义了变量, 所以需要先执行config.go
use init
main init
config.go, a = 0 b = 1 c = 2 d = 3
config.go, a = 0 b = 1 c = 2 d = 3
use.go,    a = 0 b = 1 c = 2 d = 3
main
*/
