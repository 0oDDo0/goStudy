package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*
profiling 就是对应用的画像, 即应用使用 CPU 和内存的情况
cum, 计算机利用率监视器
goroutines, go的协程
cumulative, 积累, 累加
*/

/*
Go的性能分析
CPU profile, 报告程序CPU使用情况, 按照一定频率去采集应用程序在CPU和寄存器上面的数据
Memory Profile(Heap Profile), 报告程序的内存使用情况
Block Profiling, Goroutine阻塞事件的记录
Goroutine Profiling：报告goroutines的使用情况, 有哪些goroutine, 它们的调用关系是怎样的


如果程序是用http包启动的web服务器, 想查看web服务器的状态, 可选择net/http/pprof
pprof, 会与server公用同一个端口http://127.0.0.1:8000/debug/pprof/, 以/debug/pprof/为前缀
子url
/debug/pprof/profile, 自动进行CPU profiling, 持续30s，并生成一个文件供下载
/debug/pprof/heap,    得到一个Memory Profiling结果的文件
/debug/pprof/block,   block Profiling结果
/debug/pprof/goroutine, 运行的goroutines列表以及调用关系
/debug/pprof/mutex, 查看陷入死锁的互斥锁


cmd
1. 生成CPU状态分析图, go tool pprof url/file, 默认的收集时间为30s, 可在url后加?seconds=60改为60s
	go tool pprof http://localhost:3999/debug/pprof/profile, 性能分析
	go tool pprof http://localhost:3999/debug/pprof/heap, 内存分析
	运行 go tool pprof 命令时加上 --nodefration=0.05, 表示如果调用的子函数使用的CPU,memory不超过5%,就忽略它,不要显示在图片中
	会生成pprof命令行, 根据分析文件的不同, 以下命令含义也会有些许差别
	1)  topN, 列出最耗时间的地方, 如top10, top30..., 生成如下
		flat  flat%   sum%        cum   cum%  name
		flat(%), CPU上运行的时间以及百分比;
		sum, 当前所有函数累加使用CPU的比, 即第二列= 第一列加第二列CPU使用比, 第三列等于3+2+1列...;
		cum(%), 这个函数以及子函数运行所占用的时间和比例
		name, 函数名
	2) web, 生成函数调用图
		图中每个方框对应一个函数, 方框越大代表执行的时间越久(包括它调用的子函数执行时间,但并不是正比的关系),
		方框之间的箭头代表着调用关系, 箭头上的数字代表被调用函数的执行时间.
		--inuse_objects 可以查看每个函数分配的对象数
		--alloc-space 查看分配的内存空间大小
		web funcName  只打印和funcName相关的内容

*/

var port = "3999"

func loop(name string) {
	strs := make([]string, 10*1024*1024)
	nums := []string{}
	times := 20000
	i := 0
	for {
		for j := 0; j < 1; j++ {
			nums = append(nums, "123")
		}
		if i > times && i%times == 0 {
			fmt.Println(name, i)
		}
		i++
		time.Sleep(50000) // 单位是ms
	}
	fmt.Println(strs)
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:"+port, nil))
	}()
	go loop("1")
	loop("2")
}
